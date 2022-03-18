package paint

const qryGetWallItems = `
SELECT 
	name,
	height,
	width,
	square_meters 
FROM 
	wall_items
`
