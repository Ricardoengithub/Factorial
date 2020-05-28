#lang plai
( define (factorial n) (cond 
	[(= n 0) 1]
	[else (* n (factorial (- n 1)))]))