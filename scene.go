package main

type Scene struct {
	Objects []Hitable
}

func (scene Scene) hit(r *Ray, tmin float64, tmax float64) (bool, HitRecord) {
	var finalRec, tempRecord HitRecord
	var isHit bool
	hitAnything := false
	closest := tmax
	for i := 0; i < len(scene.Objects); i++ {
		isHit, tempRecord = scene.Objects[i].hit(r, tmin, closest)
		if isHit {
			hitAnything = true
			closest = tempRecord.T
			finalRec = tempRecord
		}
	}
	return hitAnything, finalRec
}
