class main {
    public static void main(String[] args){
        StringToolsController stc = new StringToolsController();
        String name = stc.reverse("Deepak");
        String palindrome = stc.makePalindrome("Deepak");
        System.out.println(name);
        System.out.println(palindrome);
    }
}
