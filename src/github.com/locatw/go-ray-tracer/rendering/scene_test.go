package rendering

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/locatw/go-ray-tracer/element"
	"github.com/locatw/go-ray-tracer/element/mock"
	. "github.com/locatw/go-ray-tracer/image"
	. "github.com/locatw/go-ray-tracer/vector"
)

func TestLookForIntersectedObject(t *testing.T) {
	camera := CreateCamera(
		CreateZeroVector(),
		CreateAxisVector(ZAxis),
		CreateAxisVector(YAxis),
		Resolution{Width: 3, Height: 2},
		60.0,
	)
	ray := CreateRay(CreateZeroVector(), CreateAxisVector(ZAxis))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("When scene has not objects", func(t *testing.T) {
		scene := Scene{Camera: camera, Shapes: []Shape{}}

		t.Run("it returns nil", func(t *testing.T) {
			hitInfo := scene.LookForIntersectedObject(ray)

			if hitInfo != nil {
				t.Errorf("got: %v, wont: %v", hitInfo, nil)
			}
		})
	})

	t.Run("When a scene has an object which doesn't intersect a ray", func(t *testing.T) {
		shape := mock_element.NewMockShape(ctrl)
		shape.EXPECT().Intersect(ray).Return(nil)

		scene := Scene{Camera: camera, Shapes: []Shape{shape}}

		t.Run("it returns nil", func(t *testing.T) {
			hitInfo := scene.LookForIntersectedObject(ray)

			if hitInfo != nil {
				t.Errorf("got: %v, wont: %v", hitInfo, nil)
			}
		})
	})

	t.Run("When a scene has two objects which both intersects a ray", func(t *testing.T) {
		nearShape := mock_element.NewMockShape(ctrl)
		nearHitInfo := &HitInfo{
			Object:   nearShape,
			Position: Vector{X: 1.0, Y: 0.0, Z: 0.0},
			Normal:   CreateAxisVector(YAxis),
			T:        1.0,
		}
		nearShape.EXPECT().Intersect(ray).Return(nearHitInfo)

		farShape := mock_element.NewMockShape(ctrl)
		farHitInfo := &HitInfo{
			Object:   farShape,
			Position: Vector{X: 1.0, Y: 0.0, Z: 0.0},
			Normal:   CreateAxisVector(YAxis),
			T:        2.0,
		}
		farShape.EXPECT().Intersect(ray).Return(farHitInfo)

		scene := Scene{Camera: camera, Shapes: []Shape{farShape, nearShape}}

		t.Run("it returns near object", func(t *testing.T) {
			hitInfo := scene.LookForIntersectedObject(ray)

			if hitInfo != nearHitInfo {
				t.Errorf("got: %v, wont: %v", hitInfo, nearHitInfo)
			}
		})
	})
}
