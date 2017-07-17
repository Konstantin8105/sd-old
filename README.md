# GoFea

FEA for steel structural engineer on golang

**TODO`s**:

*Step 1 - Calculate truss model in 2D*

- [X] 2d point
- [X] 2b beam structure
- [X] 2d support
- [X] 2d truss finite element
- [X] 2d node load
- [X] Create global stiffiner matrix
- [X] Create load vector
- [X] Solve system of linear equation
- [X] Calculate global deformation
- [X] Calculate internal deformation
- [X] Calculate internal force
- [X] Calculate natural frequency
- [ ] Vector for linear algebra
- [ ] Add Travis checking of project
- [ ] Refactoring
- [ ] Testing
- [ ] Benchmark tests
- [ ] Release

*Step 1.1*

- [ ] Create RPC client-server for fast calculation

*Step 2*

- [ ] Buckling
- [ ] add property for truss elements
- [ ] add property for buckling analyze of case
- [ ] add natural frequency property
- [ ] New git branch
- [ ] 2d beam finite element
- [ ] 2d gravity load
- [ ] 2d displacement load
- [ ] Selfweight load
- [ ] Reactions in support
- [ ] Calculate global deformation
- [ ] Calculate internal deformation
- [ ] Calculate internal force
- [ ] Calculate natural frequency
- [ ] Buckling
- [ ] Refactoring
- [ ] Testing
- [ ] Benchmark tests
- [ ] Release

*Step 3*

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
- [ ] Refactoring
- [ ] Testing
- [ ] Benchmark tests
- [ ] Release

*Step 3.10*

- [ ] Create IO input  file format
- [ ] Create IO output file format
CALCULATION GRAPH

*Step 4*

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

*Step 5*

* MERGE MODELS
* STACK DESIGN, see https://github.com/Konstantin8105/Stack.FEA

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
