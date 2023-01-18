package fields

import "github.com/quarkcms/quark-go/pkg/component/admin/component"

type Geofence struct {
	Item
	Zoom                int    `json:"zoom"`
	MapKey              string `json:"mapKey"`
	ButtonPosition      string `json:"buttonPosition"`
	AlwaysShowItemLabel bool   `json:"alwaysShowItemLabel"`
}

// 初始化
func (p *Geofence) Init() *Geofence {
	p.Component = "geofenceField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	p.Value = map[string]interface{}{
		"center": map[string]interface{}{
			"longitude": "116.397724",
			"latitude":  "39.903755",
		},
		"points": []interface{}{},
	}
	p.Zoom = 14
	p.MapKey = "788e08def03f95c670944fe2c78fa76f"
	p.Style = map[string]interface{}{
		"height":    500,
		"width":     "100%",
		"marginTop": "10px",
	}

	return p
}

// zoom
func (p *Geofence) SetZoom(zoom int) *Geofence {
	p.Zoom = zoom
	return p
}

// 高德地图key
func (p *Geofence) SetMapKey(key string) *Geofence {
	p.MapKey = key
	return p
}

// 地图宽度
func (p *Geofence) SetWidth(width interface{}) *Geofence {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 地图高度
func (p *Geofence) SetHeight(height interface{}) *Geofence {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["height"] = height
	p.Style = style

	return p
}

// 中心点
func (p *Geofence) SetCenter(longitude string, latitude string) *Geofence {
	p.Value.(map[string]interface{})["center"] = map[string]interface{}{
		"longitude": longitude,
		"latitude":  latitude,
	}

	return p
}

// 多边形围栏坐标点
func (p *Geofence) SetPoints(points []interface{}) *Geofence {
	p.Value.(map[string]interface{})["points"] = points

	return p
}
