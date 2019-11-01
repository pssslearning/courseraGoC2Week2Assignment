# courseraGoC2Week2Assignment

## Assigment Week #2  -  FUNCTION TYPES - Course https: https://www.coursera.org/learn/golang-functions-methods/home/welcome


Tareas calificadas por los compañeros: Module 2 Activity

- [Course: Functions, Methods, and Interfaces in Go](https://www.coursera.org/learn/golang-functions-methods)

- Course 2 of 3 from `Programming with Google Go Specialization`

## Instructions

```
The goal of this assignment is to create a routine that solves a linear kinematics problem.

Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

s(t) = ½ (at²) + v0t + s0

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement. Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))

```

## My assignment

Source code at `kinematics/kinematics.go`

## A sample test compilation and execution

```sh
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC2Week2Assignment/kinematics$ cd
devel1@vbxdeb10mate:~$ 
devel1@vbxdeb10mate:~$ cd $GOPATH/src/github.com/pssslearning/courseraGoC2Week2Assignment/kinematics/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC2Week2Assignment/kinematics$ go build kinematics.go 
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC2Week2Assignment/kinematics$ ./kinematics 

------------------------------------------------------------
Please introduce the following initial configutarion values:
------------------------------------------------------------
Acceleration ......: 10
Initial Velocity...: 2
initialDisplacement: 1
Formula: displacement s(t) = ½ (10.000000 * t²) + 2.000000 * t + 1.000000
Please enter now a value for 't' (time): 5
The resulting displacement value s(t) is [136.000000]
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC2Week2Assignment/kinematics$ ./kinematics 

------------------------------------------------------------
Please introduce the following initial configutarion values:
------------------------------------------------------------
Acceleration ......: -9.8
Initial Velocity...: 100.45
initialDisplacement: -200.34
Formula: displacement s(t) = ½ (-9.800000 * t²) + 100.450000 * t + -200.340000
Please enter now a value for 't' (time): 45
The resulting displacement value s(t) is [-5602.590000]
```


