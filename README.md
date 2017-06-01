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

```
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
```

Description:
- North  - menu, toolbar, tabs
- West   - tree view for model of inlet models, tabs
- Center - 3D view of model
- East   - tables, property
- South  - status bar, 1-line with short information
> Note: 
> - flexibility border between center and west, center and east
> - tree view on West is flexibility and allowable collapse to border

North. Toobar. View options:
- 2D
- 3D view:
	- XOY
	- XOY back side
	- XOZ 
	- XOZ back side
	- YOZ
	- YOZ back side
	- XYZ(3d view)
- 3D view:
	- Simple
	- Wireframe
	- Realistic
- 3D view:
	- Orthonal
	- Perspective
	- Virtual reality

North. Toolbal. View options:
- Zoom "+"
- Zoom "-"
- Zoom all
- Zoom by window
- Hand

North toolbar. Create elements:
- New point
- New beam
- New plate
- Mirror elements
- Copy array
- Copy by circle

North toolbal. Cursor option:
- Simple cursor
- Smart cursor

North tabs (/ underline for - Center):
- Design / 3D model
- Text editor / Text
- Postprocessor / 3D model

West tabs with internal tabs (/ underline for - East), if North tab is Design:
- Overview / None
- Geometry:
	- Point / Table of points
	- Lines / Tables of points and lines
	- Plates / Tables of points and plates
- Property:
	- Shape / List of property
	- Material / List of property
	- Specific / List of property
	- Supports / List of property
	- Cases and loads / List of property
- Calculation:
	- Check code / List of property 
	- Allowable processors / List of property
	- Allowable computers / List of property

West tabs (/ underline for - East), if North tab = Postprocessor:
- Point
	- Displacement / Tabs for table displacement: cases, max/min
	- Reactions / Tabs for table reactions: cases, max/min, global
- Beam
	- Diagramm / None
	- Code ratio / None
- Plate
	- View / None

South. Label indocators:
- amount selected points
- amount points
- amount selected beams
- amount beams
- amount selected plates
- amount plates

South popup window:
- recommendations, warnings, errors


Code modification windows

MENU OPTIONS



*Step 5*

MERGE MODELS
PARALLE CALCULATION
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
