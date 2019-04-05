package main


type Material interface {
	scatter(r *Ray, rec HitRecord) (bool, Vector, Ray)
}

type Metal struct {
  Albedo Vector
}

type Lambertian struct {
  Albedo Vector
}

func (l Lambertian) scatter(r *Ray, rec HitRecord) (bool, Vector, Ray) {
  target := rec.P.Add(rec.Normal).Add(randomInUnitSphere())
  scattered := Ray{rec.P, target.Subtract(rec.P)}
  return true, l.Albedo, scattered
}

func (l Metal) scatter(r *Ray, rec HitRecord) (bool, Vector, Ray) {
  unitVector := r.Direction.MakeUnitVector()
  reflected := unitVector.Reflect(rec.Normal)
  scattered := Ray{rec.P, reflected}
  return scattered.Direction.Dot(rec.Normal) > 0 , l.Albedo, scattered
}
