package repositories

import (
	"testing"

	"goserver/internal/application/dto"
	"goserver/internal/domain/entities"
	derrors "goserver/internal/domain/errors"
)

func TestMemoryRoleRepository_CRUD(t *testing.T) {
	repo := NewMemoryRoleRepository()
	role := &entities.Role{Name: "admin", Permissions: []string{"read"}}

	// create
	created, err := repo.Create(role)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if created.ID == "" {
		t.Fatalf("expected ID to be set")
	}

	// duplicate name
	if _, err := repo.Create(&entities.Role{Name: "admin"}); err != derrors.ErrRoleAlreadyExists {
		t.Fatalf("expected duplicate error, got %v", err)
	}

	// find by id
	found, err := repo.FindByID(created.ID)
	if err != nil || found == nil || found.Name != "admin" {
		t.Fatalf("failed to find role by id")
	}

	// update
	created.Description = "super"
	updated, _ := repo.Update(created.ID, created)
	if updated.Description != "super" {
		t.Fatalf("update failed")
	}

	// pagination
	q := &dto.PaginationQuery[dto.RoleFiltersDto]{Page: 1, Limit: 10}
	res, err := repo.FindPaginated(q)
	if err != nil || len(res.Data) != 1 {
		t.Fatalf("pagination failed: %v", err)
	}

	// delete
	if _, err := repo.Delete(created.ID); err != nil {
		t.Fatalf("delete error: %v", err)
	}
	if r, _ := repo.FindByID(created.ID); r != nil {
		t.Fatalf("role should be deleted")
	}
}
