package sqlstore

type PermissionRepository struct {
	store *Store
}

func (p *PermissionRepository) CreatePermission(name string, endpoint string) error {
	return nil
}
