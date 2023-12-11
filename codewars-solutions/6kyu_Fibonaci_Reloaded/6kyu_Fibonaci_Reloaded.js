//Sample Tests
const Test = require("@codewars/test-compat");

describe("Tests", () => {
	it("test", () => {
		Test.assertEquals(fib(1), 0, "fib(1) failed");
		Test.assertEquals(fib(2), 1, "fib(2) failed");
		Test.assertEquals(fib(3), 1, "fib(3) failed");
		Test.assertEquals(fib(4), 2, "fib(4) failed");
		Test.assertEquals(fib(5), 3, "fib(5) failed");
	});
});

function fib(n) {}
