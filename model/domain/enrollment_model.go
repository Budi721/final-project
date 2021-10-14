package domain

type Enrollment struct {
    IdUser           int    `json:"id_user,omitempty"`
    NamaLengkap      string `json:"nama_lengkap,omitempty"`
    Username         string `json:"username,omitempty"`
    Password         string `json:"password,omitempty"`
    TopikDiminati    string `json:"topik_diminati,omitempty"`
    EnrollmentStatus int    `json:"enrollment_status,omitempty"`
}
