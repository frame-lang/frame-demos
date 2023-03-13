// See https://aka.ms/new-console-template for more information

using csharp.stringtools;

namespace StringTools
{
    public class Program
    {
        static void Main(string[] args)
    {
        StringToolsController stc = new StringToolsController();
        string name = stc.reverse("Deepak");
        string palindrome = stc.makePalindrome("Deepak");
        Console.WriteLine(name);
        Console.WriteLine(palindrome);
        
    }
    }
}

