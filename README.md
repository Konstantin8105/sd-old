# GoFea

[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/GoFea/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/GoFea?branch=master)
[![Build Status](https://travis-ci.org/Konstantin8105/GoFea.svg?branch=master)](https://travis-ci.org/Konstantin8105/GoFea)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/GoFea)](https://goreportcard.com/report/github.com/Konstantin8105/GoFea)

FEA for steel structural engineer on golang

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
- [x] Refactoring
- [x] Testing
- [x] Benchmark tests
- [x] Add git tag

---------------------

**TODO`s**:


*Step 0.2 - Calculate truss model in 2D*

- [x] New git branch
- [ ] Add property for nolinear deformation
- [ ] Nolinear deformation
- [ ] Reactions in support
- [ ] Add natural frequency property
- [ ] Calculate natural frequency
- [ ] Add property for buckling analyze of case
- [ ] Buckling
- [ ] Refactoring
- [ ] Testing
- [ ] Benchmark tests
- [ ] Merge pull request(PR)
- [ ] Add git tag

*Step*

- [ ] 2d beam finite element
- [ ] 2d gravity load
- [ ] 2d displacement load
- [ ] Selfweight load
- [ ] Calculate global deformation
- [ ] Calculate internal deformation
- [ ] Calculate internal force
- [ ] Calculate natural frequency
- [ ] Buckling
- [ ] Many loads for natural frequency calculation
- [ ] Redirection loads and cheching recursive loading
- [ ] New git branch
- [ ] Modal mass of frame
- [ ] Combine of truss and beams
- [ ] Pins for 2d beam finite element
- [ ] 2d uniform load
- [ ] 2d trapezoidally uniform load
- [ ] Calculate global deformation
- [ ] Calculate internal deformation
- [ ] Calculate internal force
- [ ] Calculate natural frequency
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
- [ ] Riks method for buckling
- [ ] Parallel linear algebra
- [ ] 3d node
- [ ] 3d truss finite element
- [ ] 3d node load
- [ ] rotate of shape
- [ ] axe-symmetrical finite elements + buckling. Typical case - compress load for shell
- [ ] Database of shapes
- [ ] Triangulation for user shape
- [ ] Gap finite element
- [ ] Temperature load
- [ ] Connection design
- [ ] Tension finite element
- [ ] Compress finite element
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
