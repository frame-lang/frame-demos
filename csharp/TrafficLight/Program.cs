using csharp.trafficlight;

namespace TrafficLight
{
    public class Program
    {
        static void Main(string[] args)
    {
    //     Mom mom = new Mom();
    //     mom.start();

    // System.Threading.Timer? timer = null;
    // timer = new System.Threading.Timer((obj) => {
    //     if(timer!= null){
    //     timer.Dispose();
    //     mom.stop();
    //     }
    // }, null, 10000, System.Threading.Timeout.Infinite);}
    Mom mom = new Mom();
    mom.Start();
    System.Threading.Timer? timer = new System.Threading.Timer(obj => 
    {
        mom.Stop();
    }, null, 10000, System.Threading.Timeout.Infinite);

    
    }
}
}
