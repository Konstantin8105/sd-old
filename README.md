# GoFea

[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/GoFea/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/GoFea?branch=master)
[![Build Status](https://travis-ci.org/Konstantin8105/GoFea.svg?branch=master)](https://travis-ci.org/Konstantin8105/GoFea)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/GoFea)](https://goreportcard.com/report/github.com/Konstantin8105/GoFea)

FEA for steel structural engineer on golang

Procedures of workflow:
```
+-----------------------+
| Create new git branch |
+-----------------------+
     |
	 V
+-----------+
|+---------+|
|| WORKING ||
|+---------+|
+-----------+
     |
	 V
+-------------+
| Refactoring |
+-------------+
     |
	 V
+-------------------+
| Testing, min 90 % |
+-------------------+
     |
	 V
+---------------------+
| Create pull request |
+---------------------+
     |
	 V
+---------+
| Release |
+---------+
```

#### Procedure of new calcutation type(buckling,...)
1. Add property of calcuation type in model
2. Add calculation type algorithm
3. Testing for calculation type
4. Create output of new calculation type
Tasks:
- [ ] Nolinear deformation
- [ ] Natural frequency
- [ ] Linear buckling
- [ ] Nolinear buckling
- [ ] Riks method for buckling

#### Procedure of new load type(gravity loads)
1. Create load type in model
2. Add calculation load type
3. Testing of new load type
4. Create output of new calculation type
Tasks:
- [ ] 2d gravity load (selfweight)
- [ ] 2d displacement load
- [ ] 2d uniform load
- [ ] 2d trapezoidally uniform load
- [ ] 2d temperature load
- [ ] 2d redirection of load

#### Procedure of new finite element
1. Create finite element
2. Testing finite element
3. Create output of finite element
Tasks:
- [ ] 2d beam finite element
- [ ] 2d tension finite element
- [ ] 2d compress finite element
- [ ] 2d gap finite element
- [ ] Pins for 2d beam finite element



*Step 0.1 - Calculate truss model in 2D*

- [x] 2d point
- [x] 2d support
- [x] 2d truss finite element
- [x] 2d node load
- [x] Create global stiffiner matrix
- [x] Create load vector
- [x] Solve system of linear equation
- [x] Calculate global deformation
- [x] Calculate internal deformation
- [x] Calculate internal force

---------------------

**TODO`s**:


*Step 0.2 - Calculate truss model in 2D*

- [ ] Reactions in support
- [ ] Many loads for natural frequency calculation
- [ ] Redirection loads and cheching recursive loading
- [ ] Create IO input  file format
- [ ] Create IO output file format
- [ ] CALCULATION GRAPH
- [ ] Minimal build configuration. Best way is Go native code.
- [ ] 3D, tables, graphs
- [ ] HTML+CSS gui
- [ ] Design of GUI for 2D elements

threejs
https://threejs.org/
http://davidscottlyons.com/threejs/presentations/frontporch14/#slide-0
https://threejs.org/editor/
https://github.com/mrdoob/three.js/tree/master/editor
maybe - polymer from google

- [ ] MERGE MODELS
- [ ] STACK DESIGN, see https://github.com/Konstantin8105/Stack.FEA
- [ ] add checking - point cannot have same number
- [ ] Create RPC client-server for fast calculation
- [ ] working with geometry inside
- [ ] triangulation of regions
- [ ] Intersection between beams
- [ ] Separate finite elements on small elements
- [ ] Parallel linear algebra
- [ ] rotate of shape
- [ ] axe-symmetrical finite elements + buckling. Typical case - compress load for shell
- [ ] Database of shapes
- [ ] Triangulation for user shape
- [ ] Connection design
- [ ] Sparse matrix solver
- [ ] Time dependence
- [ ] Nonlinear property of material. Temperature, corrosion.
- [ ] Load force in global system coordinate
- [ ] Load force in local system coordinate - important for non-linear buckling. [Ko] * Zo + [Go] * Zo = Po and Po is dependence of local point rotation.
- [ ] Tests for all parts
- [ ] Test buckling on tension !!! Important
- [ ] Code review
- [ ] Code coverage
- [ ] Add examples of calculation
- [ ] Contribution rules
- [ ] Automatic tool for truss frames
- [ ] Where is money?
