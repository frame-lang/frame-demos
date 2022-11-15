package javaframedemos.PersistentTrafficLight;
import java.util.concurrent.TimeUnit;

public class Main{
    public static void main(String[] args) throws Exception {
        TrafficLightMomController mom = new TrafficLightMomController();
        int count=0;
        while(count<=3){
            System.out.println("Tick");
            TimeUnit.SECONDS.sleep(1);
            mom.tick();
            count++;
        }
        System.out.println("Finished");                 
    }
    
}