public class main {
    public static void main(String[] args) {
        HistoryBasicController basic = new HistoryBasicController();
        basic.start();
        basic.switchState();
        basic.gotoDeadEnd();
        basic.back();

        HistoryStateContextController compartment = new HistoryStateContextController()
        compartment.start();
        compartment.switchState();
        compartment.gotoDeadEnd();
        compartment.back();
    }
}