/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package cmd

import (
	"fmt"
	"net"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mysteriumnetwork/node/core/connection/connectionstate"
	"github.com/mysteriumnetwork/node/money"

	"github.com/mysteriumnetwork/node/communication/nats"
	"github.com/mysteriumnetwork/node/config"
	appconfig "github.com/mysteriumnetwork/node/config"
	"github.com/mysteriumnetwork/node/consumer/bandwidth"
	consumer_session "github.com/mysteriumnetwork/node/consumer/session"
	"github.com/mysteriumnetwork/node/consumer/statistics"
	"github.com/mysteriumnetwork/node/core/auth"
	"github.com/mysteriumnetwork/node/core/connection"
	"github.com/mysteriumnetwork/node/core/discovery/brokerdiscovery"
	"github.com/mysteriumnetwork/node/core/discovery/proposal"
	"github.com/mysteriumnetwork/node/core/ip"
	"github.com/mysteriumnetwork/node/core/location"
	"github.com/mysteriumnetwork/node/core/node"
	nodevent "github.com/mysteriumnetwork/node/core/node/event"
	"github.com/mysteriumnetwork/node/core/policy"
	"github.com/mysteriumnetwork/node/core/port"
	"github.com/mysteriumnetwork/node/core/quality"
	"github.com/mysteriumnetwork/node/core/service"
	"github.com/mysteriumnetwork/node/core/state"
	"github.com/mysteriumnetwork/node/core/storage/boltdb"
	"github.com/mysteriumnetwork/node/core/storage/boltdb/migrations/history"
	"github.com/mysteriumnetwork/node/core/storage/boltdb/migrator"
	"github.com/mysteriumnetwork/node/eventbus"
	"github.com/mysteriumnetwork/node/feedback"
	"github.com/mysteriumnetwork/node/firewall"
	"github.com/mysteriumnetwork/node/identity"
	"github.com/mysteriumnetwork/node/identity/registry"
	identity_registry "github.com/mysteriumnetwork/node/identity/registry"
	identity_selector "github.com/mysteriumnetwork/node/identity/selector"
	"github.com/mysteriumnetwork/node/logconfig"
	"github.com/mysteriumnetwork/node/market/mysterium"
	"github.com/mysteriumnetwork/node/metadata"
	"github.com/mysteriumnetwork/node/mmn"
	"github.com/mysteriumnetwork/node/nat"
	"github.com/mysteriumnetwork/node/nat/event"
	"github.com/mysteriumnetwork/node/nat/mapping"
	"github.com/mysteriumnetwork/node/nat/traversal"
	"github.com/mysteriumnetwork/node/nat/upnp"
	"github.com/mysteriumnetwork/node/p2p"
	"github.com/mysteriumnetwork/node/requests"
	"github.com/mysteriumnetwork/node/services"
	service_noop "github.com/mysteriumnetwork/node/services/noop"
	service_openvpn "github.com/mysteriumnetwork/node/services/openvpn"
	"github.com/mysteriumnetwork/node/session/connectivity"
	"github.com/mysteriumnetwork/node/session/pingpong"
	"github.com/mysteriumnetwork/node/sleep"
	"github.com/mysteriumnetwork/node/tequilapi"
	tequilapi_endpoints "github.com/mysteriumnetwork/node/tequilapi/endpoints"
	"github.com/mysteriumnetwork/node/utils"
	"github.com/mysteriumnetwork/node/utils/netutil"

	paymentClient "github.com/mysteriumnetwork/payments/client"
	"github.com/mysteriumnetwork/payments/uniswap"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// UIServer represents our web server
type UIServer interface {
	Serve() error
	Stop()
}

// Dependencies is DI container for top level components which is reused in several places
type Dependencies struct {
	Node *Node

	HTTPClient *requests.HTTPClient

	NetworkDefinition metadata.NetworkDefinition
	MysteriumAPI      *mysterium.MysteriumAPI
	EtherClient       *paymentClient.ReconnectableEthClient
	Exchange          *money.Exchange

	BrokerConnector  *nats.BrokerConnector
	BrokerConnection nats.Connection

	NATService       nat.NATService
	Storage          *boltdb.Bolt
	Keystore         *identity.Keystore
	IdentityManager  identity.Manager
	SignerFactory    identity.SignerFactory
	IdentityRegistry identity_registry.IdentityRegistry
	IdentitySelector identity_selector.Handler

	DiscoveryFactory   service.DiscoveryFactory
	ProposalRepository proposal.Repository
	DiscoveryWorker    brokerdiscovery.Worker

	QualityClient *quality.MysteriumMORQA

	IPResolver       ip.Resolver
	LocationResolver *location.Cache

	PolicyOracle *policy.Oracle

	StatisticsReporter               *statistics.SessionStatisticsReporter
	SessionStorage                   *consumer_session.Storage
	SessionConnectivityStatusStorage connectivity.StatusStorage

	EventBus eventbus.EventBus

	ConnectionManager  connection.Manager
	ConnectionRegistry *connection.Registry

	ServicesManager *service.Manager
	ServiceRegistry *service.Registry
	ServiceSessions *service.SessionPool
	ServiceFirewall firewall.IncomingTrafficFirewall

	NATPinger  traversal.NATPinger
	NATTracker *event.Tracker
	PortPool   *port.Pool
	PortMapper mapping.PortMapper

	StateKeeper *state.Keeper

	P2PDialer   p2p.Dialer
	P2PListener p2p.Listener

	Authenticator     *auth.Authenticator
	JWTAuthenticator  *auth.JWTAuthenticator
	UIServer          UIServer
	Transactor        *registry.Transactor
	BCHelper          *paymentClient.BlockchainWithRetries
	ProviderRegistrar *registry.ProviderRegistrar

	LogCollector *logconfig.Collector
	Reporter     *feedback.Reporter

	ProviderInvoiceStorage   *pingpong.ProviderInvoiceStorage
	ConsumerTotalsStorage    *pingpong.ConsumerTotalsStorage
	HermesPromiseStorage     *pingpong.HermesPromiseStorage
	ConsumerBalanceTracker   *pingpong.ConsumerBalanceTracker
	HermesChannelRepository  *pingpong.HermesChannelRepository
	HermesPromiseSettler     pingpong.HermesPromiseSettler
	HermesURLGetter          *pingpong.HermesURLGetter
	HermesCaller             *pingpong.HermesCaller
	ChannelAddressCalculator *pingpong.ChannelAddressCalculator
	HermesPromiseHandler     *pingpong.HermesPromiseHandler
	SettlementHistoryStorage *pingpong.SettlementHistoryStorage

	MMN *mmn.MMN
}

// Bootstrap initiates all container dependencies
func (di *Dependencies) Bootstrap(nodeOptions node.Options) error {
	logconfig.Configure(&nodeOptions.LogOptions)

	netutil.LogNetworkStats()

	p2p.RegisterContactUnserializer()
	di.BrokerConnector = nats.NewBrokerConnector()

	log.Info().Msg("Starting Mysterium Node " + metadata.VersionAsString())

	di.HTTPClient = requests.NewHTTPClient(nodeOptions.BindAddress, requests.DefaultTimeout)

	// Check early for presence of an already running node
	tequilaListener, err := di.createTequilaListener(nodeOptions)
	if err != nil {
		return err
	}

	if err := nodeOptions.Directories.Check(); err != nil {
		return err
	}

	if err := di.bootstrapFirewall(nodeOptions.Firewall); err != nil {
		return err
	}

	di.bootstrapEventBus()

	if err := di.bootstrapStorage(nodeOptions.Directories.Storage); err != nil {
		return err
	}

	netutil.ClearStaleRoutes()

	if err := di.bootstrapNetworkComponents(nodeOptions); err != nil {
		return err
	}

	di.bootstrapIdentityComponents(nodeOptions)

	if err := di.bootstrapDiscoveryComponents(nodeOptions.Discovery); err != nil {
		return err
	}
	if err := di.bootstrapLocationComponents(nodeOptions); err != nil {
		return err
	}
	if err := di.bootstrapAuthenticator(); err != nil {
		return err
	}

	di.bootstrapUIServer(nodeOptions)
	if err := di.bootstrapMMN(); err != nil {
		return err
	}
	if err := di.bootstrapNATComponents(nodeOptions); err != nil {
		return err
	}

	di.PortPool = port.NewPool()
	if config.GetBool(config.FlagPortMapping) {
		portmapConfig := mapping.DefaultConfig()
		di.PortMapper = mapping.NewPortMapper(portmapConfig, di.EventBus)
	} else {
		di.PortMapper = mapping.NewNoopPortMapper(di.EventBus)
	}

	di.bootstrapP2P(nodeOptions.P2PPorts)
	di.SessionConnectivityStatusStorage = connectivity.NewStatusStorage()

	if err := di.bootstrapServices(nodeOptions); err != nil {
		return err
	}

	if err := di.bootstrapQualityComponents(nodeOptions.BindAddress, nodeOptions.Quality); err != nil {
		return err
	}

	if err := di.bootstrapNodeComponents(nodeOptions, tequilaListener); err != nil {
		return err
	}

	di.registerConnections(nodeOptions)
	if err = di.handleConnStateChange(); err != nil {
		return err
	}
	if err := di.Node.Start(); err != nil {
		return err
	}

	appconfig.Current.EnableEventPublishing(di.EventBus)

	log.Info().Msg("Mysterium node started!")
	return nil
}

func (di *Dependencies) bootstrapP2P(p2pPorts *port.Range) {
	portPool := di.PortPool
	natPinger := di.NATPinger
	identityVerifier := identity.NewVerifierSigned()
	if p2pPorts.IsSpecified() {
		log.Info().Msgf("Fixed p2p service port range (%s) configured, using custom port pool", p2pPorts)
		portPool = port.NewFixedRangePool(*p2pPorts)
		natPinger = traversal.NewNoopPinger()
	}

	di.P2PListener = p2p.NewListener(di.BrokerConnection, di.SignerFactory, identityVerifier, di.IPResolver, natPinger, portPool, di.PortMapper)
	di.P2PDialer = p2p.NewDialer(di.BrokerConnector, di.SignerFactory, identityVerifier, di.IPResolver, natPinger, portPool)
}

func (di *Dependencies) createTequilaListener(nodeOptions node.Options) (net.Listener, error) {
	if !nodeOptions.TequilapiEnabled {
		return tequilapi.NewNoopListener()
	}

	tequilaListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", nodeOptions.TequilapiAddress, nodeOptions.TequilapiPort))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("the port %v seems to be taken. Either you're already running a node or it is already used by another application", nodeOptions.TequilapiPort))
	}
	return tequilaListener, nil
}

func (di *Dependencies) bootstrapStateKeeper(options node.Options) error {
	var lastStageName string
	if options.ExperimentNATPunching {
		lastStageName = traversal.StageName
	} else {
		lastStageName = mapping.StageName
	}

	deps := state.KeeperDeps{
		NATStatusProvider:         nat.NewStatusTracker(lastStageName),
		Publisher:                 di.EventBus,
		ServiceLister:             di.ServicesManager,
		IdentityProvider:          di.IdentityManager,
		IdentityRegistry:          di.IdentityRegistry,
		IdentityChannelCalculator: di.ChannelAddressCalculator,
		BalanceProvider:           di.ConsumerBalanceTracker,
		EarningsProvider:          di.HermesChannelRepository,
	}
	di.StateKeeper = state.NewKeeper(deps, state.DefaultDebounceDuration)
	return di.StateKeeper.Subscribe(di.EventBus)
}

func (di *Dependencies) registerOpenvpnConnection(nodeOptions node.Options) {
	service_openvpn.Bootstrap()
	connectionFactory := func() (connection.Connection, error) {
		return service_openvpn.NewClient(
			// TODO instead of passing binary path here, Openvpn from node options could represent abstract vpn factory itself
			nodeOptions.Openvpn.BinaryPath(),
			nodeOptions.Directories.Script,
			nodeOptions.Directories.Runtime,
			di.SignerFactory,
			di.IPResolver,
		)
	}
	di.ConnectionRegistry.Register(service_openvpn.ServiceType, connectionFactory)
}

func (di *Dependencies) registerNoopConnection() {
	service_noop.Bootstrap()
	di.ConnectionRegistry.Register(service_noop.ServiceType, service_noop.NewConnection)
}

// Shutdown stops container
func (di *Dependencies) Shutdown() (err error) {
	var errs []error
	defer func() {
		for i := range errs {
			log.Error().Err(errs[i]).Msg("Dependencies shutdown failed")
			if err == nil {
				err = errs[i]
			}
		}
	}()

	// Kill node first which includes current active VPN connection cleanup.
	if di.Node != nil {
		if err := di.Node.Kill(); err != nil {
			errs = append(errs, err)
		}
	}

	if di.ServicesManager != nil {
		if err := di.ServicesManager.Kill(); err != nil {
			errs = append(errs, err)
		}
	}

	if di.PolicyOracle != nil {
		di.PolicyOracle.Stop()
	}

	if di.NATService != nil {
		if err := di.NATService.Disable(); err != nil {
			errs = append(errs, err)
		}
	}
	if di.DiscoveryWorker != nil {
		di.DiscoveryWorker.Stop()
	}
	if di.BrokerConnection != nil {
		di.BrokerConnection.Close()
	}

	if di.QualityClient != nil {
		di.QualityClient.Stop()
	}

	if di.ServiceFirewall != nil {
		di.ServiceFirewall.Teardown()
	}
	firewall.Reset()

	if di.Storage != nil {
		if err := di.Storage.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	return nil
}

func (di *Dependencies) bootstrapStorage(path string) error {
	localStorage, err := boltdb.NewStorage(path)
	if err != nil {
		return err
	}

	migrator := migrator.NewMigrator(localStorage)
	err = migrator.RunMigrations(history.Sequence)
	if err != nil {
		return err
	}

	di.Storage = localStorage

	if !config.GetBool(config.FlagUserMode) {
		netutil.SetRouteManagerStorage(di.Storage)
	}

	invoiceStorage := pingpong.NewInvoiceStorage(di.Storage)
	di.ProviderInvoiceStorage = pingpong.NewProviderInvoiceStorage(invoiceStorage)
	di.ConsumerTotalsStorage = pingpong.NewConsumerTotalsStorage(di.Storage, di.EventBus)
	di.HermesPromiseStorage = pingpong.NewHermesPromiseStorage(di.Storage)
	di.SessionStorage = consumer_session.NewSessionStorage(di.Storage)
	di.SettlementHistoryStorage = pingpong.NewSettlementHistoryStorage(di.Storage)
	return di.SessionStorage.Subscribe(di.EventBus)
}

func (di *Dependencies) bootstrapNodeComponents(nodeOptions node.Options, tequilaListener net.Listener) error {
	// Consumer current session bandwidth
	bandwidthTracker := bandwidth.NewTracker(di.EventBus)
	if err := bandwidthTracker.Subscribe(di.EventBus); err != nil {
		return err
	}

	// Consumer session history (API storage)
	di.StatisticsReporter = statistics.NewSessionStatisticsReporter(di.MysteriumAPI, di.SignerFactory, time.Minute)
	if err := di.StatisticsReporter.Subscribe(di.EventBus); err != nil {
		return err
	}

	di.Transactor = registry.NewTransactor(
		di.HTTPClient,
		nodeOptions.Transactor.TransactorEndpointAddress,
		nodeOptions.Transactor.RegistryAddress,
		nodeOptions.Hermes.HermesID,
		nodeOptions.Transactor.ChannelImplementation,
		di.SignerFactory,
		di.EventBus,
		di.BCHelper,
	)

	if err := di.bootstrapHermesPromiseSettler(nodeOptions); err != nil {
		return err
	}

	if err := di.bootstrapProviderRegistrar(nodeOptions); err != nil {
		return err
	}

	di.ChannelAddressCalculator = pingpong.NewChannelAddressCalculator(
		nodeOptions.Hermes.HermesID,
		nodeOptions.Transactor.ChannelImplementation,
		nodeOptions.Transactor.RegistryAddress,
	)

	hermesURL, err := di.HermesURLGetter.GetHermesURL(common.HexToAddress(nodeOptions.Hermes.HermesID))
	if err != nil {
		return err
	}

	di.HermesCaller = pingpong.NewHermesCaller(di.HTTPClient, hermesURL)
	di.ConsumerBalanceTracker = pingpong.NewConsumerBalanceTracker(
		di.EventBus,
		common.HexToAddress(nodeOptions.Payments.MystSCAddress),
		common.HexToAddress(nodeOptions.Hermes.HermesID),
		di.BCHelper,
		di.ChannelAddressCalculator,
		di.ConsumerTotalsStorage,
		di.HermesCaller,
		di.Transactor,
		di.IdentityRegistry,
	)

	err = di.ConsumerBalanceTracker.Subscribe(di.EventBus)
	if err != nil {
		return errors.Wrap(err, "could not subscribe consumer balance tracker to relevant events")
	}

	di.HermesPromiseHandler = pingpong.NewHermesPromiseHandler(pingpong.HermesPromiseHandlerDeps{
		HermesPromiseStorage: di.HermesPromiseStorage,
		HermesCallerFactory: func(hermesURL string) pingpong.HermesHTTPRequester {
			return pingpong.NewHermesCaller(di.HTTPClient, hermesURL)
		},
		HermesURLGetter: di.HermesURLGetter,
		FeeProvider:     di.Transactor,
		Encryption:      di.Keystore,
		EventBus:        di.EventBus,
	})

	if err := di.HermesPromiseHandler.Subscribe(di.EventBus); err != nil {
		return err
	}

	di.ConnectionRegistry = connection.NewRegistry()
	di.ConnectionManager = connection.NewManager(
		pingpong.ExchangeFactoryFunc(
			di.Keystore,
			di.SignerFactory,
			di.ConsumerTotalsStorage,
			nodeOptions.Transactor.ChannelImplementation,
			nodeOptions.Transactor.RegistryAddress,
			di.EventBus,
			nodeOptions.Payments.ConsumerDataLeewayMegabytes,
		),
		di.ConnectionRegistry.CreateConnection,
		di.EventBus,
		di.IPResolver,
		di.LocationResolver,
		connection.DefaultConfig(),
		connection.DefaultStatsReportInterval,
		connection.NewValidator(
			di.ConsumerBalanceTracker,
			di.IdentityManager,
		),
		di.P2PDialer,
	)

	di.LogCollector = logconfig.NewCollector(&logconfig.CurrentLogOptions)
	reporter, err := feedback.NewReporter(di.LogCollector, di.IdentityManager, nodeOptions.FeedbackURL)
	if err != nil {
		return err
	}
	di.Reporter = reporter

	if err := di.bootstrapStateKeeper(nodeOptions); err != nil {
		return err
	}

	uniswapClient := money.NewUniswapClient(func(c *ethclient.Client) *uniswap.Client {
		return uniswap.NewClient(c)
	}, di.EtherClient)
	di.Exchange = money.NewExchange(
		common.HexToAddress(nodeOptions.Payments.MystSCAddress),
		common.HexToAddress(nodeOptions.Payments.DaiAddress),
		common.HexToAddress(nodeOptions.Payments.WethAddress),
		uniswapClient,
	)

	tequilapiHTTPServer, err := di.bootstrapTequilapi(nodeOptions, tequilaListener)
	if err != nil {
		return err
	}

	sleepNotifier := sleep.NewNotifier(di.ConnectionManager, di.EventBus)
	sleepNotifier.Subscribe()

	di.Node = NewNode(di.ConnectionManager, tequilapiHTTPServer, di.EventBus, di.NATPinger, di.UIServer, sleepNotifier)
	return nil
}

func (di *Dependencies) bootstrapTequilapi(nodeOptions node.Options, listener net.Listener) (tequilapi.APIServer, error) {
	if !nodeOptions.TequilapiEnabled {
		return tequilapi.NewNoopAPIServer(), nil
	}

	router := tequilapi.NewAPIRouter()
	tequilapi_endpoints.AddRoutesForDocs(router)
	tequilapi_endpoints.AddRouteForStop(router, utils.SoftKiller(di.Shutdown))
	tequilapi_endpoints.AddRoutesForAuthentication(router, di.Authenticator, di.JWTAuthenticator)
	tequilapi_endpoints.AddRoutesForIdentities(router, di.IdentityManager, di.IdentitySelector, di.IdentityRegistry, di.ConsumerBalanceTracker, di.ChannelAddressCalculator, di.HermesChannelRepository, di.BCHelper, di.Transactor)
	tequilapi_endpoints.AddRoutesForConnection(router, di.ConnectionManager, di.StateKeeper, di.ProposalRepository, di.IdentityRegistry)
	tequilapi_endpoints.AddRoutesForSessions(router, di.SessionStorage)
	tequilapi_endpoints.AddRoutesForConnectionLocation(router, di.IPResolver, di.LocationResolver, di.LocationResolver)
	tequilapi_endpoints.AddRoutesForProposals(router, di.ProposalRepository, di.QualityClient)
	tequilapi_endpoints.AddRoutesForService(router, di.ServicesManager, services.JSONParsersByType)
	tequilapi_endpoints.AddRoutesForPayout(router, di.IdentityManager, di.SignerFactory, di.MysteriumAPI)
	tequilapi_endpoints.AddRoutesForAccessPolicies(di.HTTPClient, router, config.GetString(config.FlagAccessPolicyAddress))
	tequilapi_endpoints.AddRoutesForNAT(router, di.StateKeeper)
	tequilapi_endpoints.AddRoutesForTransactor(router, di.Transactor, di.HermesPromiseSettler, di.SettlementHistoryStorage, common.HexToAddress(nodeOptions.Hermes.HermesID))
	tequilapi_endpoints.AddRoutesForConfig(router)
	tequilapi_endpoints.AddRoutesForMMN(router, di.MMN)
	tequilapi_endpoints.AddRoutesForFeedback(router, di.Reporter)
	tequilapi_endpoints.AddRoutesForConnectivityStatus(router, di.SessionConnectivityStatusStorage)
	tequilapi_endpoints.AddRoutesForCurrencyExchange(router, di.Exchange)
	if err := tequilapi_endpoints.AddRoutesForSSE(router, di.StateKeeper, di.EventBus); err != nil {
		return nil, err
	}

	if config.GetBool(config.FlagPProfEnable) {
		tequilapi_endpoints.AddRoutesForPProf(router)
	}

	corsPolicy := tequilapi.NewMysteriumCorsPolicy()
	return tequilapi.NewServer(listener, router, corsPolicy), nil
}

// function decides on network definition combined from testnet/localnet flags and possible overrides
func (di *Dependencies) bootstrapNetworkComponents(options node.Options) (err error) {
	optionsNetwork := options.OptionsNetwork
	network := metadata.DefaultNetwork

	switch {
	case optionsNetwork.Betanet:
		network = metadata.BetanetDefinition
	case optionsNetwork.Testnet:
		network = metadata.TestnetDefinition
	case optionsNetwork.Localnet:
		network = metadata.LocalnetDefinition
	}

	//override defined values one by one from options
	if optionsNetwork.MysteriumAPIAddress != metadata.DefaultNetwork.MysteriumAPIAddress {
		network.MysteriumAPIAddress = optionsNetwork.MysteriumAPIAddress
	}

	if optionsNetwork.BrokerAddress != metadata.DefaultNetwork.BrokerAddress {
		network.BrokerAddress = optionsNetwork.BrokerAddress
	}

	if optionsNetwork.EtherClientRPC != metadata.DefaultNetwork.EtherClientRPC {
		network.EtherClientRPC = optionsNetwork.EtherClientRPC
	}

	di.NetworkDefinition = network

	di.MysteriumAPI = mysterium.NewClient(di.HTTPClient, network.MysteriumAPIAddress)

	brokerURL, err := nats.ParseServerURI(di.NetworkDefinition.BrokerAddress)
	if err != nil {
		return err
	}
	if _, err := di.ServiceFirewall.AllowURLAccess(brokerURL.String()); err != nil {
		return err
	}
	if di.BrokerConnection, err = di.BrokerConnector.Connect(brokerURL.String()); err != nil {
		return err
	}

	log.Info().Msg("Using Eth endpoint: " + network.EtherClientRPC)
	di.EtherClient, err = paymentClient.NewReconnectableEthClient(network.EtherClientRPC)
	if err != nil {
		return err
	}

	bc := paymentClient.NewBlockchain(di.EtherClient, options.Payments.BCTimeout)
	di.BCHelper = paymentClient.NewBlockchainWithRetries(bc, time.Millisecond*300, 3)

	di.HermesURLGetter = pingpong.NewHermesURLGetter(di.BCHelper, common.HexToAddress(options.Transactor.RegistryAddress))

	registryStorage := registry.NewRegistrationStatusStorage(di.Storage)
	if di.IdentityRegistry, err = identity_registry.NewIdentityRegistryContract(di.EtherClient, common.HexToAddress(options.Transactor.RegistryAddress), common.HexToAddress(options.Hermes.HermesID), registryStorage, di.EventBus); err != nil {
		return err
	}

	hermesURL, err := di.HermesURLGetter.GetHermesURL(common.HexToAddress(options.Hermes.HermesID))
	if err != nil {
		return err
	}

	if _, err := firewall.AllowURLAccess(
		network.EtherClientRPC,
		network.MysteriumAPIAddress,
		options.Transactor.TransactorEndpointAddress,
		hermesURL,
	); err != nil {
		return err
	}
	if _, err := di.ServiceFirewall.AllowURLAccess(
		network.EtherClientRPC,
		network.MysteriumAPIAddress,
		options.Transactor.TransactorEndpointAddress,
		hermesURL,
	); err != nil {
		return err
	}

	return di.IdentityRegistry.Subscribe(di.EventBus)
}

func (di *Dependencies) bootstrapEventBus() {
	di.EventBus = eventbus.New()
}

func (di *Dependencies) bootstrapIdentityComponents(options node.Options) {
	var ks *keystore.KeyStore
	if options.Keystore.UseLightweight {
		log.Debug().Msg("Using lightweight keystore")
		ks = keystore.NewKeyStore(options.Directories.Keystore, keystore.LightScryptN, keystore.LightScryptP)
	} else {
		log.Debug().Msg("Using heavyweight keystore")
		ks = keystore.NewKeyStore(options.Directories.Keystore, keystore.StandardScryptN, keystore.StandardScryptP)
	}

	di.Keystore = identity.NewKeystoreFilesystem(options.Directories.Keystore, ks)
	di.IdentityManager = identity.NewIdentityManager(di.Keystore, di.EventBus)
	di.SignerFactory = func(id identity.Identity) identity.Signer {
		return identity.NewSigner(di.Keystore, id)
	}
	di.IdentitySelector = identity_selector.NewHandler(
		di.IdentityManager,
		di.MysteriumAPI,
		identity.NewIdentityCache(options.Directories.Keystore, "remember.json"),
		di.SignerFactory,
	)

}

func (di *Dependencies) bootstrapQualityComponents(bindAddress string, options node.OptionsQuality) (err error) {
	if _, err := firewall.AllowURLAccess(options.Address); err != nil {
		return err
	}
	if _, err := di.ServiceFirewall.AllowURLAccess(options.Address); err != nil {
		return err
	}
	di.QualityClient = quality.NewMorqaClient(bindAddress, options.Address, di.SignerFactory, 10*time.Second)
	go di.QualityClient.Start()

	var transport quality.Transport
	switch options.Type {
	case node.QualityTypeElastic:
		transport = quality.NewElasticSearchTransport(di.HTTPClient, options.Address, 10*time.Second)
	case node.QualityTypeMORQA:
		transport = quality.NewMORQATransport(di.QualityClient)
	case node.QualityTypeNone:
		transport = quality.NewNoopTransport()
	default:
		err = errors.Errorf("unknown Quality Oracle provider: %s", options.Type)
	}
	if err != nil {
		return err
	}

	fn := di.LocationResolver.RunnerCountryFn()
	// Quality metrics
	qualitySender := quality.NewSender(transport, metadata.VersionAsString(), fn)
	if err := qualitySender.Subscribe(di.EventBus); err != nil {
		return err
	}

	// warm up the loader as the load takes up to a couple of secs
	loader := &upnp.GatewayLoader{}
	go loader.Get()
	natSender := event.NewSender(qualitySender, di.IPResolver.GetPublicIP, loader.HumanReadable)
	if err := natSender.Subscribe(di.EventBus); err != nil {
		return err
	}

	return nil
}

func (di *Dependencies) bootstrapLocationComponents(options node.Options) (err error) {
	if _, err = firewall.AllowURLAccess(options.Location.IPDetectorURL); err != nil {
		return errors.Wrap(err, "failed to add firewall exception")
	}
	if _, err = di.ServiceFirewall.AllowURLAccess(options.Location.IPDetectorURL); err != nil {
		return errors.Wrap(err, "failed to add firewall exception")
	}
	ipResolver := ip.NewResolver(di.HTTPClient, options.BindAddress, options.Location.IPDetectorURL)
	di.IPResolver = ip.NewCachedResolver(ipResolver, 5*time.Minute)

	var resolver location.Resolver
	switch options.Location.Type {
	case node.LocationTypeManual:
		resolver = location.NewStaticResolver(options.Location.Country, options.Location.City, options.Location.NodeType, di.IPResolver)
	case node.LocationTypeBuiltin:
		resolver, err = location.NewBuiltInResolver(di.IPResolver)
	case node.LocationTypeMMDB:
		resolver, err = location.NewExternalDBResolver(filepath.Join(options.Directories.Script, options.Location.Address), di.IPResolver)
	case node.LocationTypeOracle:
		if _, err := firewall.AllowURLAccess(options.Location.Address); err != nil {
			return err
		}
		if _, err := di.ServiceFirewall.AllowURLAccess(options.Location.Address); err != nil {
			return err
		}
		resolver, err = location.NewOracleResolver(di.HTTPClient, options.Location.Address), nil
	default:
		err = errors.Errorf("unknown location provider: %s", options.Location.Type)
	}
	if err != nil {
		return err
	}

	di.LocationResolver = location.NewCache(resolver, time.Minute*5)

	err = di.EventBus.SubscribeAsync(connectionstate.AppTopicConnectionState, di.LocationResolver.HandleConnectionEvent)
	if err != nil {
		return err
	}

	err = di.EventBus.SubscribeAsync(nodevent.AppTopicNode, di.LocationResolver.HandleNodeEvent)
	if err != nil {
		return err
	}

	return nil
}

func (di *Dependencies) bootstrapAuthenticator() error {
	key, err := auth.NewJWTEncryptionKey(di.Storage)
	if err != nil {
		return err
	}
	di.Authenticator = auth.NewAuthenticator(di.Storage)
	di.JWTAuthenticator = auth.NewJWTAuthenticator(key)

	return nil
}

func (di *Dependencies) bootstrapNATComponents(options node.Options) error {
	di.NATTracker = event.NewTracker()
	if err := di.NATTracker.Subscribe(di.EventBus); err != nil {
		return err
	}

	if options.ExperimentNATPunching {
		log.Debug().Msg("Experimental NAT punching enabled, creating a pinger")
		di.NATPinger = traversal.NewPinger(
			traversal.DefaultPingConfig(),
			di.EventBus,
		)
	} else {
		di.NATPinger = &traversal.NoopPinger{}
	}
	return nil
}

func (di *Dependencies) bootstrapFirewall(options node.OptionsFirewall) error {
	firewall.DefaultOutgoingFirewall = firewall.NewOutgoingTrafficFirewall(config.GetBool(config.FlagOutgoingFirewall))
	if err := firewall.DefaultOutgoingFirewall.Setup(); err != nil {
		return err
	}

	di.ServiceFirewall = firewall.NewIncomingTrafficFirewall(config.GetBool(config.FlagIncomingFirewall))
	if err := di.ServiceFirewall.Setup(); err != nil {
		return err
	}

	if options.BlockAlways {
		bindAddress := "0.0.0.0"
		resolver := ip.NewResolver(di.HTTPClient, bindAddress, "")
		outboundIP, err := resolver.GetOutboundIP()
		if err != nil {
			return err
		}

		_, err = firewall.BlockNonTunnelTraffic(firewall.Global, outboundIP)
		return err
	}
	return nil
}

func (di *Dependencies) handleConnStateChange() error {
	if di.HTTPClient == nil {
		return errors.New("HTTPClient is not initialized")
	}

	latestState := connectionstate.NotConnected
	return di.EventBus.SubscribeAsync(connectionstate.AppTopicConnectionState, func(e connectionstate.AppEventConnectionState) {
		// Here we care only about connected and disconnected events.
		if e.State != connectionstate.Connected && e.State != connectionstate.NotConnected {
			return
		}

		isDisconnected := latestState == connectionstate.Connected && e.State == connectionstate.NotConnected
		isConnected := latestState == connectionstate.NotConnected && e.State == connectionstate.Connected
		if isDisconnected || isConnected {
			netutil.LogNetworkStats()

			log.Info().Msg("Reconnecting HTTP clients due to VPN connection state change")
			di.HTTPClient.Reconnect()
			di.QualityClient.Reconnect()

			if err := di.EtherClient.Reconnect(); err != nil {
				log.Warn().Err(err).Msg("Ethereum client failed to reconnect, will retry one more time")
				// Default golang DNS resolver does not allow to reload /etc/resolv.conf more than once per 5 seconds.
				// This could lead to the problem, when right after connect/disconnect new DNS config not applied instantly.
				// Doing a couple of retries here to make sure we reconnected Ethererum client correctly.
				// Default DNS timeout is 10 seconds. It's enough to try to reconnect only twice to cover 5 seconds lag for DNS config reload.
				// https://github.com/mysteriumnetwork/node/issues/2282
				if err := di.EtherClient.Reconnect(); err != nil {
					log.Error().Err(err).Msg("Ethereum client failed to reconnect")
				}
			}
			di.EventBus.Publish(registry.AppTopicEthereumClientReconnected, struct{}{})
		}
		latestState = e.State
	})
}
