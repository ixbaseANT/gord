package main
import (
	"github.com/pkg/errors"
	"db"
)

func main() {
    err := db.InitConnection()
    if err != nil {
        panic(err)
    }
	subCmd, config := parseCommandLine()

	var err error
	switch subCmd {
	case createSubCmd:
		err = create(config.(*createConfig))
	case balanceSubCmd:
		err = balance(config.(*balanceConfig))
	case sendSubCmd:
		err = send(config.(*sendConfig))
	case createUnsignedTransactionSubCmd:
		err = createUnsignedTransaction(config.(*createUnsignedTransactionConfig))
	case signSubCmd:
		err = sign(config.(*signConfig))
	case broadcastSubCmd:
		err = broadcast(config.(*broadcastConfig))
	case parseSubCmd:
		err = parse(config.(*parseConfig))
	case showAddressesSubCmd:
		err = showAddresses(config.(*showAddressesConfig))
	case newAddressSubCmd:
		err = newAddress(config.(*newAddressConfig))
	case dumpUnencryptedDataSubCmd:
		err = dumpUnencryptedData(config.(*dumpUnencryptedDataConfig))
	case startDaemonSubCmd:
		err = startDaemon(config.(*startDaemonConfig))
	case sweepSubCmd:
		err = sweep(config.(*sweepConfig))
	default:
		err = errors.Errorf("Unknown sub-command '%s'\n", subCmd)
	}

	if err != nil {
		printErrorAndExit(err)
	}
}
