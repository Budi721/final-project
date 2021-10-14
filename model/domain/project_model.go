package domain

type Project struct {
    IdProject          int    `json:"id_project,omitempty" gorm:"primaryKey"`
    KategoriProject    string `json:"kategori_project,omitempty"`
    NamaProject        string `json:"nama_project,omitempty"`
    TanggalMulai       int64  `json:"tanggal_mulai,omitempty"`
    LinkTrello         string `json:"link_trello,omitempty"`
    DeskripsiProject   string `json:"deskripsi_project,omitempty"`
    InvitedUserId      []int  `json:"invited_user_id,omitempty" gorm:"-"`
    CollaboratorUserId []int  `json:"collaborator_user_id,omitempty" gorm:"-"`
    Admin              int    `json:"admin,omitempty"`
    UsersInvited       []User `json:"-" gorm:"many2many:user_invited;"`
    UsersCollaborator  []User `json:"-" gorm:"many2many:user_collaborator;"`
}
