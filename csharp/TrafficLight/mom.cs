using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using csharp.trafficlight;

namespace csharp.trafficlight
{
    

class Mom
{
    private TrafficLightController mom;

    public Mom()
    {
        this.mom = new TrafficLightController();
    }

    public void Start()
    {
        this.mom.start(this.mom);
    }

    public void Stop()
    {
        this.mom.stop();
    }
}

}