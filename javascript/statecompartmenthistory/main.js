const HistoryStateContextController = require("./stateCompartmentHistory");
const HistoryBasicController = require("./StateHistoryBasic");


const main = () => {

    const basic = new HistoryBasicController()
    basic.start()
    basic.switchState()
	basic.gotoDeadEnd()
	basic.back()

    const compartment = new HistoryStateContextController()
	compartment.start()
	compartment.switchState()
	compartment.gotoDeadEnd()
	compartment.back()

}

main()