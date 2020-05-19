def factorial(n):
    if(n == 1 or n == 0):
        return 1
    else:
        return n * factorial(n-1)
    

def main():
    
    n = int(input("Enter a number: "))
    
    if n >= 0:
        print(f"Factorial: {factorial(n)}")
    else:
        print(f"Choose another number")
  

if __name__ == "__main__":
    main()
