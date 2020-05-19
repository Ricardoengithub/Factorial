def factorial(n):
    if(n == 0 or n == 1):
        return 1
    fact = 1
    for i in range(2,n+1):
        fact*= i
    return fact
    

def main():
    
    n = int(input("Enter a number: "))
    
    if n >= 0:
        print(f"Factorial: {factorial(n)}")
    else:
        print(f"Choose another number")
  

if __name__ == "__main__":
    main()
