const StringToolsController = require("./machine");

const tools = new StringToolsController();

const name = tools.reverse("MARK");

console.log(name);
const palindrome = tools.makePalindrome(name);
console.log(palindrome);
