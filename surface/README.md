# Surface
## Surface computes an SVG rendering of a 3D surface function
This program demonstrates the mapping between 3 different coordinate systems
1. 2D grid of 100x100 cells identified by interger coordinates (i,j)
2. Mesh of 3D floating point coordinates (x,y,z) where x and y are linear functions of i and j, starting at (0,0) plotting from the back to the front so that background polygons may be obscured by foreground ones.
3. 2D image canvas at (0,0) points in this plane are denoted by (sx,sy) an isometric projection to map each 3D point (x,y,z) onto the 2D canvas