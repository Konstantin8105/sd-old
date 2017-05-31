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
- [ ] Buckling

*Step 2*

- [ ] 2d beam finite element
- [ ] 2d gravity load
- [ ] 2d displacement load
- [ ] Reactions in support
- [ ] Calculate global deformation
- [ ] Calculate internal deformation
- [ ] Calculate internal force
- [ ] Calculate natural frequency
- [ ] Buckling

*Step 3*

- [ ] Combine of truss and beams
- [ ] Pins for 2d beam finite element
- [ ] 2d uniform load
- [ ] 2d trapezoidally uniform load
- [ ] Calculate global deformation
- [ ] Calculate internal deformation
- [ ] Calculate internal force
- [ ] Calculate natural frequency

*Step 3.10*

- [ ] Create IO input  file format
- [ ] Create IO output file format
CALCULATION GRAPH

*Step 4*

- [ ] HTML+CSS gui
- [ ] Design of GUI for 2D elements

threejs
https://threejs.org/
http://davidscottlyons.com/threejs/presentations/frontporch14/#slide-0
https://threejs.org/editor/
https://github.com/mrdoob/three.js/tree/master/editor

maybe - polymer from google

**GUI**

'''
Main window:
+----------------------------------------+
|             North                      |
+----------------------------------------+
|         |                    |         |
| West    |    Center          | East    |
|         |                    |         |
|         |                    |         |
|         |                    |         |
|         |                    |         |
|         |                    |         |
+----------------------------------------+
|             South                      |
+----------------------------------------+
'''

Description:
North  - menu, toolbar, tabs
West   - tree view for model, innet models, tabs
Center - 3D view of model
East   - tables, property
South  - status bar, 1-line with short information

View options:
- 2D
- 3D: 6 sides, 3d view

Design tabs:
- Overview
- Geometry:
	- Point
	- Lines
	- Plates
- Property:
	- Shape
	- Material
	- Specific
	- Supports
	- Cases and loads
- Calculation:
	- Check code
	- Allowable processors
	- Allowable computers

Postprocessor view
- Point
	- Displacement
	- Reactions
- Beam
	- Dia
	- Code ratio
- Plate
	- View

Code modification windows

MENU OPTIONS



*Step 5*

PARALLE
 CALCULATION
 STACK DESIGN
- [ ] Riks method for buckling
- [ ] Parallel linear algebra
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
