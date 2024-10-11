package tracer

import (
	"cmp"
	"slices"
)

type BvhNode struct {
	left, right Hittable
	bbox        Aabb
}

func NewBvhNode(objects HittableList) BvhNode {
	var left Hittable
	var right Hittable
	if len(objects) == 1 {
		left = objects[0]
		right = objects[0]
	} else if len(objects) == 2 {
		left = objects[0]
		right = objects[1]
	} else {
		slices.SortFunc(objects, func(a, b Hittable) int {
			return cmp.Compare(a.BoundingBox().X.Min, b.BoundingBox().X.Min)
		})
		l := make(HittableList, 0)
		for i := range len(objects) / 2 {
			l = append(l, objects[i])
		}
		r := make(HittableList, 0)
		for i := len(objects) / 2; i < len(objects); i++ {
			r = append(r, objects[i])
		}
		right = &r
	}

	bbox := NewBoundingBoxFromBoundingBoxes(left.BoundingBox(), right.BoundingBox())

	return BvhNode{
		left:  left,
		right: right,
		bbox:  bbox,
	}
}

func (bvh *BvhNode) Hit(ray Ray, interval Interval, rec *HitRecord) bool {
	if !bvh.bbox.Hit(ray, interval, rec) {
		return false
	}

	hitLeft := bvh.left.Hit(ray, interval, rec)
	hitRightInterval := NewInterval(interval.Min, interval.Max)
	if hitLeft {
		hitRightInterval.Max = rec.T
	}
	hitRight := bvh.right.Hit(ray, hitRightInterval, rec)

	return hitLeft || hitRight
}

func (bvh *BvhNode) BoundingBox() Aabb {
	return bvh.bbox
}
