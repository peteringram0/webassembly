// declare function sayHello(): void;

// sayHello();

export function add(x: i32, y: i32): i32 {
  return x + y;
}

export function fib(num: i32): i32 {
  if (num <= 1) return 1;
  return fib(num - 1) + fib(num - 2);
}