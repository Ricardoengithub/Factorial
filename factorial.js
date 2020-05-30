function factorial(f) {
  // eslint-disable-next-line
  if(f == 1 || f == 0){
    return 1
  }
  if(f < 0){
    return "Error"
  }
  return f * factorial(f-1) 
}

console.log(factorial(4))
