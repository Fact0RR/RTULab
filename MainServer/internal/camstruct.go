package internal

type Camera struct{
	Id string `json:"id"`
	Type string `json:"type"`
	Xcoordinate float64 `json:"xcoordinate"`
	Ycoordinate float64 `json:"ycoordinate"`
	Description string `json:"description"`
}
// create table if not exists cameras(
// 	id varchar(255) primary key,
// 	type varchar(255) not null,
// 	coordinateX real not null,
// 	coordinateY real not null,
// 	description text 
//  );

 