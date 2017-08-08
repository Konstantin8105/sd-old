# GoFea

[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/GoFea/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/GoFea?branch=master)
[![Build Status](https://travis-ci.org/Konstantin8105/GoFea.svg?branch=master)](https://travis-ci.org/Konstantin8105/GoFea)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/GoFea)](https://goreportcard.com/report/github.com/Konstantin8105/GoFea)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Konstantin8105/GoFea/blob/master/LICENSE)

FEA for steel structural engineer on golang

Some options:
* Minimal build configuration
* Best way is Go native code

### Procedures of workflow
1. Create new git branch
2. **WORKING**
3. Refactoring
4. Testing, min 90 %
5. Create pull request
6. Release

### Procedure of new calcutation type(buckling,...)
1. Add property of calcuation type in model
2. Add calculation type algorithm
3. Testing for calculation type
4. Create output of new calculation type
Tasks:
- Linear deformation
	- Create load vector
	- Create global stiffiner matrix
	- Solve system of linear equation
	- Calculate global deformation
	- Calculate internal deformation
	- Calculate internal force
	- Reactions in support
- Nolinear deformation
- Natural frequency
- Linear buckling
- Nolinear buckling
- Riks method for buckling

### Procedure of new load type(gravity loads)
1. Create load type in model
2. Add calculation load type
3. Testing of new load type
4. Create output of new calculation type
Tasks:
- node load
- gravity load (selfweight)
- displacement load
- uniform load. local axe
- uniform load. global axe
- trapezoidally uniform load
- temperature load
- redirection of load, cheching recursive

### Procedure of new finite element
1. Create finite element
2. Testing finite element
3. Create output of finite element
Tasks:
- 2d truss finite element
- 2d beam finite element
- 2d tension finite element
- 2d compress finite element
- 2d gap finite element
- Pins for 2d beam finite element

### Procedure of new dimension
- 2d
	- 2d point
	- 2d support
- 2d symmetric
- 3d

---------------------

- [ ] 2d
	- [x] point
	- [x] support
		- [x] truss finite element
			- [x] node load
				- [x] Linear deformation
					- [x] Create load vector
					- [x] Create global stiffiner matrix
					- [x] Solve system of linear equation
					- [x] Calculate global deformation
					- [x] Calculate internal deformation
					- [x] Calculate internal force
					- [x] Reactions in support
				- [x] Natural frequency
				- [ ] Nolinear deformation
				- [ ] Nolinear deformation shapes
				- [x] Linear buckling
				- [ ] Linear buckling shapes
				- [ ] Nolinear buckling
				- [ ] Nolinear buckling shapes
				- [ ] Riks method for buckling
				- [ ] Riks method for buckling shapes
				- [ ] Natural frequency. Modal mass
				- [ ] Natural frequency shapes
			- [ ] gravity load (selfweight)
			- [ ] displacement load
			- [ ] uniform load. local axe
			- [ ] uniform load. global axe
			- [ ] trapezoidally uniform load
			- [ ] temperature load
			- [ ] redirection of load, cheching recursive
		- [ ] beam finite element
		- [ ] tension finite element
		- [ ] compress finite element
		- [ ] gap finite element
		- [ ] Pins for beam finite element
		- [ ] triangle finite element
- [ ] 2d symmetric
- [ ] 3d
- [ ] Create IO input  file format
- [ ] Create IO output file format
- [ ] HTML+CSS gui

---------------------

Tests for finite elements:

|       | Linear deformation | Nolinear deformation | Natural frequency |Linear buckling | Nolinear buckling |
| ----- |:------------------:|:--------------------:|:-----------------:|:--------------:|:-----------------:|
| Algorithm | Done           |                      | Done              | Done           |                   |
| Truss | Done               |                      | Done              |                |                   |
| Beam  | Done               |                      | Done              | Done           |                   |

---------------------

**TODO`s**:

- [ ] Fuzzer for good model with different loads(positive, negative)
- [ ] Fuzzer for change position of part code for good model
- [ ] Generate testing
- [ ] Calculation graph
- [ ] 3D, tables, graphs
- [ ] UI design https://github.com/Konstantin8105/GoFeaGUI
* threejs https://threejs.org/
* http://davidscottlyons.com/threejs/presentations/frontporch14/#slide-0
* https://threejs.org/editor/
* https://github.com/mrdoob/three.js/tree/master/editor
* maybe - polymer from google
* http://www.qooxdoo.org/

- [ ] Merge models
- [ ] Stack design, see https://github.com/Konstantin8105/Stack.FEA
- [ ] Connection design

- [x] add checking - point cannot have same number
- [ ] Create RPC client-server for fast calculation
- [ ] working with geometry inside
- [ ] Parallel linear algebra
- [ ] Geometry block for geometry operations:
	- [ ] rotate of shape
	- [ ] Intersection between beams
	- [ ] triangulation of regions
	- [ ] rotate the geometry
	- [ ] add element, point, ...
	- [ ] add forces...
	- [ ] create combinations
	- [ ] Separate finite elements on small elements
- [ ] Database of shapes
- [ ] Triangulation for user shape
- [ ] Sparse matrix solver
	- [ ] Benchmark tests
- [ ] Nonlinear property of material. Temperature, corrosion.
- [ ] Load force in local system coordinate - important for non-linear buckling. [Ko] * Zo + [Go] * Zo = Po and Po is dependence of local point rotation.
- [ ] Test buckling on tension !!! Important
- [ ] Add examples of calculation
- [ ] Contribution rules
- [ ] Parametric models
	- [ ] Automatic tool for truss frames
