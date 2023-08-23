package entity

type Criminal_report struct {
	ID            int    `json:"id"`
	Hero_id       string `json:"hero_id"`
	Villain_id    string `json:"villain_id"`
	Description   string `json:"description"`
	Incident_time string `json:"incident_time"`
}

// CREATE TABLE criminal_reports (
//     id INT AUTO_INCREMENT PRIMARY KEY,
//     hero_id INT,
//     villain_id INT,
//     description VARCHAR(255),
//     incident_time DATETIME,
//     FOREIGN KEY (hero_id) REFERENCES heroes(id),
//     FOREIGN KEY (villain_id) REFERENCES villains(id)
// );