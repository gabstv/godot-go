package godot

import (
	"github.com/gabstv/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewAStarFromPointer(ptr gdnative.Pointer) AStar {
func newAStarFromPointer(ptr gdnative.Pointer) AStar {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AStar{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A* (A star) is a computer algorithm that is widely used in pathfinding and graph traversal, the process of plotting short paths among vertices (points), passing through a given set of edges (segments). It enjoys widespread use due to its performance and accuracy. Godot's A* implementation uses points in three-dimensional space and Euclidean distances by default. You must add points manually with [method add_point] and create segments manually with [method connect_points]. Then you can test if there is a path between two points with the [method are_points_connected] function, get a path containing indices by [method get_id_path], or one containing actual coordinates with [method get_point_path]. It is also possible to use non-Euclidean distances. To do so, create a class that extends [code]AStar[/code] and override methods [method _compute_cost] and [method _estimate_cost]. Both take two indices and return a length, as is shown in the following example. [codeblock] class MyAStar: extends AStar func _compute_cost(u, v): return abs(u - v) func _estimate_cost(u, v): return min(0, abs(u - v) - 1) [/codeblock] [method _estimate_cost] should return a lower bound of the distance, i.e. [code]_estimate_cost(u, v) <= _compute_cost(u, v)[/code]. This serves as a hint to the algorithm because the custom [code]_compute_cost[/code] might be computation-heavy. If this is not the case, make [method _estimate_cost] return the same value as [method _compute_cost] to provide the algorithm with the most accurate information.
*/
type AStar struct {
	Reference
	owner gdnative.Object
}

func (o *AStar) BaseClass() string {
	return "AStar"
}

/*
        Called when computing the cost between two connected points. Note that this function is hidden in the default [code]AStar[/code] class.
	Args: [{ false from_id int} { false to_id int}], Returns: float
*/
func (o *AStar) X_ComputeCost(fromId gdnative.Int, toId gdnative.Int) gdnative.Real {
	//log.Println("Calling AStar.X_ComputeCost()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(fromId)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "_compute_cost")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Called when estimating the cost between a point and the path's ending point. Note that this function is hidden in the default [code]AStar[/code] class.
	Args: [{ false from_id int} { false to_id int}], Returns: float
*/
func (o *AStar) X_EstimateCost(fromId gdnative.Int, toId gdnative.Int) gdnative.Real {
	//log.Println("Calling AStar.X_EstimateCost()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(fromId)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "_estimate_cost")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Adds a new point at the given position with the given identifier. The algorithm prefers points with lower [code]weight_scale[/code] to form a path. The [code]id[/code] must be 0 or larger, and the [code]weight_scale[/code] must be 1 or larger. [codeblock] var astar = AStar.new() astar.add_point(1, Vector3(1, 0, 0), 4) # Adds the point (1, 0, 0) with weight_scale 4 and id 1 [/codeblock] If there already exists a point for the given [code]id[/code], its position and weight scale are updated to the given values.
	Args: [{ false id int} { false position Vector3} {1 true weight_scale float}], Returns: void
*/
func (o *AStar) AddPoint(id gdnative.Int, position gdnative.Vector3, weightScale gdnative.Real) {
	//log.Println("Calling AStar.AddPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromVector3(position)
	ptrArguments[2] = gdnative.NewPointerFromReal(weightScale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "add_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns whether the two given points are directly connected by a segment. If [code]bidirectional[/code] is [code]false[/code], returns whether movement from [code]id[/code] to [code]to_id[/code] is possible through this segment.
	Args: [{ false id int} { false to_id int} {True true bidirectional bool}], Returns: bool
*/
func (o *AStar) ArePointsConnected(id gdnative.Int, toId gdnative.Int, bidirectional gdnative.Bool) gdnative.Bool {
	//log.Println("Calling AStar.ArePointsConnected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)
	ptrArguments[2] = gdnative.NewPointerFromBool(bidirectional)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "are_points_connected")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Clears all the points and segments.
	Args: [], Returns: void
*/
func (o *AStar) Clear() {
	//log.Println("Calling AStar.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Creates a segment between the given points. If [code]bidirectional[/code] is [code]false[/code], only movement from [code]id[/code] to [code]to_id[/code] is allowed, not the reverse direction. [codeblock] var astar = AStar.new() astar.add_point(1, Vector3(1, 1, 0)) astar.add_point(2, Vector3(0, 5, 0)) astar.connect_points(1, 2, false) [/codeblock]
	Args: [{ false id int} { false to_id int} {True true bidirectional bool}], Returns: void
*/
func (o *AStar) ConnectPoints(id gdnative.Int, toId gdnative.Int, bidirectional gdnative.Bool) {
	//log.Println("Calling AStar.ConnectPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)
	ptrArguments[2] = gdnative.NewPointerFromBool(bidirectional)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "connect_points")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Deletes the segment between the given points. If [code]bidirectional[/code] is [code]false[/code], only movement from [code]id[/code] to [code]to_id[/code] is prevented, and a unidirectional segment possibly remains.
	Args: [{ false id int} { false to_id int} {True true bidirectional bool}], Returns: void
*/
func (o *AStar) DisconnectPoints(id gdnative.Int, toId gdnative.Int, bidirectional gdnative.Bool) {
	//log.Println("Calling AStar.DisconnectPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)
	ptrArguments[2] = gdnative.NewPointerFromBool(bidirectional)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "disconnect_points")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the next available point ID with no point associated to it.
	Args: [], Returns: int
*/
func (o *AStar) GetAvailablePointId() gdnative.Int {
	//log.Println("Calling AStar.GetAvailablePointId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "get_available_point_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the ID of the closest point to [code]to_position[/code], optionally taking disabled points into account. Returns [code]-1[/code] if there are no points in the points pool. [b]Note:[/b] If several points are the closest to [code]to_position[/code], the one with the smallest ID will be returned, ensuring a deterministic result.
	Args: [{ false to_position Vector3} {False true include_disabled bool}], Returns: int
*/
func (o *AStar) GetClosestPoint(toPosition gdnative.Vector3, includeDisabled gdnative.Bool) gdnative.Int {
	//log.Println("Calling AStar.GetClosestPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector3(toPosition)
	ptrArguments[1] = gdnative.NewPointerFromBool(includeDisabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "get_closest_point")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the closest position to [code]to_position[/code] that resides inside a segment between two connected points. [codeblock] var astar = AStar.new() astar.add_point(1, Vector3(0, 0, 0)) astar.add_point(2, Vector3(0, 5, 0)) astar.connect_points(1, 2) var res = astar.get_closest_position_in_segment(Vector3(3, 3, 0)) # Returns (0, 3, 0) [/codeblock] The result is in the segment that goes from [code]y = 0[/code] to [code]y = 5[/code]. It's the closest position in the segment to the given point.
	Args: [{ false to_position Vector3}], Returns: Vector3
*/
func (o *AStar) GetClosestPositionInSegment(toPosition gdnative.Vector3) gdnative.Vector3 {
	//log.Println("Calling AStar.GetClosestPositionInSegment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(toPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "get_closest_position_in_segment")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns an array with the IDs of the points that form the path found by AStar between the given points. The array is ordered from the starting point to the ending point of the path. [codeblock] var astar = AStar.new() astar.add_point(1, Vector3(0, 0, 0)) astar.add_point(2, Vector3(0, 1, 0), 1) # Default weight is 1 astar.add_point(3, Vector3(1, 1, 0)) astar.add_point(4, Vector3(2, 0, 0)) astar.connect_points(1, 2, false) astar.connect_points(2, 3, false) astar.connect_points(4, 3, false) astar.connect_points(1, 4, false) var res = astar.get_id_path(1, 3) # Returns [1, 2, 3] [/codeblock] If you change the 2nd point's weight to 3, then the result will be [code][1, 4, 3][/code] instead, because now even though the distance is longer, it's "easier" to get through point 4 than through point 2.
	Args: [{ false from_id int} { false to_id int}], Returns: PoolIntArray
*/
func (o *AStar) GetIdPath(fromId gdnative.Int, toId gdnative.Int) gdnative.PoolIntArray {
	//log.Println("Calling AStar.GetIdPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(fromId)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "get_id_path")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the capacity of the structure backing the points, useful in conjunction with [code]reserve_space[/code].
	Args: [], Returns: int
*/
func (o *AStar) GetPointCapacity() gdnative.Int {
	//log.Println("Calling AStar.GetPointCapacity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "get_point_capacity")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns an array with the IDs of the points that form the connection with the given point. [codeblock] var astar = AStar.new() astar.add_point(1, Vector3(0, 0, 0)) astar.add_point(2, Vector3(0, 1, 0)) astar.add_point(3, Vector3(1, 1, 0)) astar.add_point(4, Vector3(2, 0, 0)) astar.connect_points(1, 2, true) astar.connect_points(1, 3, true) var neighbors = astar.get_point_connections(1) # Returns [2, 3] [/codeblock]
	Args: [{ false id int}], Returns: PoolIntArray
*/
func (o *AStar) GetPointConnections(id gdnative.Int) gdnative.PoolIntArray {
	//log.Println("Calling AStar.GetPointConnections()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "get_point_connections")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the number of points currently in the points pool.
	Args: [], Returns: int
*/
func (o *AStar) GetPointCount() gdnative.Int {
	//log.Println("Calling AStar.GetPointCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "get_point_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns an array with the points that are in the path found by AStar between the given points. The array is ordered from the starting point to the ending point of the path.
	Args: [{ false from_id int} { false to_id int}], Returns: PoolVector3Array
*/
func (o *AStar) GetPointPath(fromId gdnative.Int, toId gdnative.Int) gdnative.PoolVector3Array {
	//log.Println("Calling AStar.GetPointPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(fromId)
	ptrArguments[1] = gdnative.NewPointerFromInt(toId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "get_point_path")

	// Call the parent method.
	// PoolVector3Array
	retPtr := gdnative.NewEmptyPoolVector3Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector3ArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the position of the point associated with the given [code]id[/code].
	Args: [{ false id int}], Returns: Vector3
*/
func (o *AStar) GetPointPosition(id gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling AStar.GetPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "get_point_position")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the weight scale of the point associated with the given [code]id[/code].
	Args: [{ false id int}], Returns: float
*/
func (o *AStar) GetPointWeightScale(id gdnative.Int) gdnative.Real {
	//log.Println("Calling AStar.GetPointWeightScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "get_point_weight_scale")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns an array of all points.
	Args: [], Returns: Array
*/
func (o *AStar) GetPoints() gdnative.Array {
	//log.Println("Calling AStar.GetPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "get_points")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Returns whether a point associated with the given [code]id[/code] exists.
	Args: [{ false id int}], Returns: bool
*/
func (o *AStar) HasPoint(id gdnative.Int) gdnative.Bool {
	//log.Println("Calling AStar.HasPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "has_point")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns whether a point is disabled or not for pathfinding. By default, all points are enabled.
	Args: [{ false id int}], Returns: bool
*/
func (o *AStar) IsPointDisabled(id gdnative.Int) gdnative.Bool {
	//log.Println("Calling AStar.IsPointDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "is_point_disabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Removes the point associated with the given [code]id[/code] from the points pool.
	Args: [{ false id int}], Returns: void
*/
func (o *AStar) RemovePoint(id gdnative.Int) {
	//log.Println("Calling AStar.RemovePoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "remove_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Reserves space internally for [code]num_nodes[/code] points, useful if you're adding a known large number of points at once, for a grid for instance. New capacity must be greater or equals to old capacity.
	Args: [{ false num_nodes int}], Returns: void
*/
func (o *AStar) ReserveSpace(numNodes gdnative.Int) {
	//log.Println("Calling AStar.ReserveSpace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(numNodes)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "reserve_space")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Disables or enables the specified point for pathfinding. Useful for making a temporary obstacle.
	Args: [{ false id int} {True true disabled bool}], Returns: void
*/
func (o *AStar) SetPointDisabled(id gdnative.Int, disabled gdnative.Bool) {
	//log.Println("Calling AStar.SetPointDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromBool(disabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "set_point_disabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the [code]position[/code] for the point with the given [code]id[/code].
	Args: [{ false id int} { false position Vector3}], Returns: void
*/
func (o *AStar) SetPointPosition(id gdnative.Int, position gdnative.Vector3) {
	//log.Println("Calling AStar.SetPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromVector3(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "set_point_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the [code]weight_scale[/code] for the point with the given [code]id[/code].
	Args: [{ false id int} { false weight_scale float}], Returns: void
*/
func (o *AStar) SetPointWeightScale(id gdnative.Int, weightScale gdnative.Real) {
	//log.Println("Calling AStar.SetPointWeightScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromReal(weightScale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AStar", "set_point_weight_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AStarImplementer is an interface that implements the methods
// of the AStar class.
type AStarImplementer interface {
	ReferenceImplementer
	X_ComputeCost(fromId gdnative.Int, toId gdnative.Int) gdnative.Real
	X_EstimateCost(fromId gdnative.Int, toId gdnative.Int) gdnative.Real
	AddPoint(id gdnative.Int, position gdnative.Vector3, weightScale gdnative.Real)
	ArePointsConnected(id gdnative.Int, toId gdnative.Int, bidirectional gdnative.Bool) gdnative.Bool
	Clear()
	ConnectPoints(id gdnative.Int, toId gdnative.Int, bidirectional gdnative.Bool)
	DisconnectPoints(id gdnative.Int, toId gdnative.Int, bidirectional gdnative.Bool)
	GetAvailablePointId() gdnative.Int
	GetClosestPoint(toPosition gdnative.Vector3, includeDisabled gdnative.Bool) gdnative.Int
	GetClosestPositionInSegment(toPosition gdnative.Vector3) gdnative.Vector3
	GetIdPath(fromId gdnative.Int, toId gdnative.Int) gdnative.PoolIntArray
	GetPointCapacity() gdnative.Int
	GetPointConnections(id gdnative.Int) gdnative.PoolIntArray
	GetPointCount() gdnative.Int
	GetPointPath(fromId gdnative.Int, toId gdnative.Int) gdnative.PoolVector3Array
	GetPointPosition(id gdnative.Int) gdnative.Vector3
	GetPointWeightScale(id gdnative.Int) gdnative.Real
	GetPoints() gdnative.Array
	HasPoint(id gdnative.Int) gdnative.Bool
	IsPointDisabled(id gdnative.Int) gdnative.Bool
	RemovePoint(id gdnative.Int)
	ReserveSpace(numNodes gdnative.Int)
	SetPointDisabled(id gdnative.Int, disabled gdnative.Bool)
	SetPointPosition(id gdnative.Int, position gdnative.Vector3)
	SetPointWeightScale(id gdnative.Int, weightScale gdnative.Real)
}
