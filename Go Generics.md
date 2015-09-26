### Generics in Go

The last significant thread on generics in the Go users forum ended largely on the question of, well, what would generics in Go look like? A pretty wide open question. But quite necessary as a first step in identifying the requirements, consequences and complications of any implementaion. 

Here then is one possible answer.

#### Defined

First, what does generics in Go actually mean.

All common definitions of generics leap directly into a deep, set-theoretic discussion of type systems, extended type relations and axiomatic derivations.  Certainly appropriate where a generics systems is intended to handle exotics, such as 'temporary inferred borrows' between objects in a dynamic, associative hierarchy.

Not really helpful.  So, back to first principles.

The most basic definition of generics is simply that it is an abstraction that enables type-agnostic implementations of algorithms. Implicitly, set-theoretics are not fundamental to generics. 

This provides a good starting point, since the design of Go itself is not based on set theory.

To consider what this definition could comprehend in the context of Go requires applying or extending the core design principles of Go in general to the definition. The relevant design  features of Go include type-safety, minimalism, and clarity of expression. In addition, the Go Authors have stated or stongly suggested that
 
* generics are desired
* generics should blend seamlessly with the existing language features
* language changes can not break backwards compatibility
* any required language additions and enhancements must be minimal and well-targted

The best realistic definition of generics in Go is then simply 

         a minimal abstraction mechanism that enables the type-safe, 
             type-agnostic implementation of algorithms with a clean  
             and efficient presentation and use. 

#### Design

Consider an idealized generic implementation of an algorithm, say computation of a dot product of two vectors V of degree-n, returning a scalar result S:

    func Dot_V(a, b V) (res S) {
    	 d := a.LenMin_V(b)
    	 for p := 0; pos < d; p++ {
	    	 i := a.Get_V(p)
    		 j := b.Get_V(p)
	     	res.Add_S(i.Mul_S(j))
    	 }
	    return res
    }

A straight-forward statement of the algorithm. Just needs type-specialization and maybe something to make use clearer.

Turns out, type specialization only requires an implementation strategy. Specifically, that V and S are defined by method-set interfaces that specify the minimum set of atomic functions necessary to externalize type-specific operations.

    type V interface {
    	Get_V(int) S
    	Add_V(int, V)
    	MulSc_V(int, S)
    	. . . .
      LenMin_V(V) int
    }

    type S interface {
	    Add_S(S) S
	    Mul_S(S) S
	    . . . .
    }

That, coupled with name-spacing the type-specific implementations in distinct packages. Obviously.
 
Ease-of-use can be achived by adding light-weight – facade like – type-specific functions. Again, in distinct, conveniently name-spaced packages. Calls use type actuals and returns hide the type assertions. The resulting use API presents no appearence of generic implementation.

    package floats
    
    func (a *Vector32) Dot32(b *Vector32) float32 {
    	 return float32(Dot_V(a, b).(Scalar32))
    }
    
    package big
    
    func (a *Vector) Dot(b *Vector) big.Float {
       return big.Float(Dot_V(a, b).(Scalar))
}
    

Collectively, a Go generics pattern.

#### Proof-of-Concept Implementations

An initial vector maths proof-of-concept implementation, even including exemplary generic linear and spherical vector interpolation algorithms, is available at github.com/grosenberg/go-generics.  Float32 and big.Float type specializations are provided.

Generic collections and rich enums, implemented using the same pattern, are also provided.

#### Issues

The point of developing this possible solution was to begin identifying and evaluating the issues that must be faced in this and similar generics implementation in Go.

These issues, in no particular order, include

- the excessive length required to implement at least some of the type-specialization functions. From the POV of the generic routine, each should be a high-level, yet atomic or nearly atomic function. In practice, the functions require several, if not 10's of lines to implement. Ideally, commonalities should be identified and considered for some form of generic implementation. As yet, no conclusions can be drawn as to the optimal set of these language feature changes and enhancements.  

- the complexity of the reflection work sometimes required. This is particularly true in the collections implemention. It is noisy in terms of presentation at least, particularly when required in the type-agnostic routines. The utlities collected in the generic package are intended to minimize this noise. Optimizing the selection and signatures of these utilities will help. Additional reflection runtime support may be necessary, but no conclusions can be drawn at this time.

- the asymetric absence of an interface demotion operator for values passed on function calls. Parameter and return values are nominally promoted to interface values when passed to an interface parameter or assignment variable. In order to return a generic value in assignment to a typed variable, an interface demotion operator language feature addition is required. For purposes of discusion, the interface demotion operator is denoted as `.()` forces an implicit reflect.ValueOf() on the value being returned as it is passed. All existing rules concerning the assignment of a value to a variable are unaffected, including failing if reflect.ValueOf(r) cannot be automatically assigned to `x`. 

               type R interface{}
               var x, y someType

               // type-specific call of generic function
               x = someFunction(y) 

               // generic function with generic to type-specific return
               func someFunction(r R) R.() { 
                   return r
               }

- the absence of a true type-specialization notation. This would be implemented as a language feature enhancement with the goal of implicitly providing the equivalent utility of the separate user API. Implementation of the interface demotion operator would be required to meet this goal.

- the inability to declare variables and functions private to a limited set of packages or otherwise declare name-spaces. While the user API needs to be public, the type-specific implementations should be only be internally visible. But that requires visiblity across the internal package boundaries without being public to all external packages. Any number of minor language feature enhancements could achive this goal. Potentially moot if true type-specialization is implemented, depending on the nature of the implementation.

- the potential need to explicitly verify type-safety. Since interface types are potentially fungible, checks may have to be implemented to ensure intended results. Vigilance is required.

- the potential complexity to extend to higher-orders of generic elements -- generics in generics. Go favors compositional constructions. The complexity of implementing higher-order generics is, however, unknown. It is believed that on the order of 80% of actual use cases only use first-order generics, that a substantial majority of the remainder can be handed by appropriately designed first-order generics, and that higher-order generics can be constructed with modest effort, where required. This belief needs to be tested.    


