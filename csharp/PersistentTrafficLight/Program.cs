using csharp.persistenttrafficlight;

namespace PersistentTrafficLight
{
    public class Program
    {
        static void Main(string[] args)
    {
    
    TrafficLightMomController mom = new TrafficLightMomController();
        int count=0;
        while(count<=3){
            Console.WriteLine("Tick");
            mom.tick();
            count++;
        }
        Console.WriteLine("Finished");                 
    }

        
    
    }
}

