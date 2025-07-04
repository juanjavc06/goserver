package services

import (
	"testing"

	"goserver/internal/application/dto"
	"goserver/internal/infrastructure/repositories"
)

func TestRoleService_CreateAndFind(t *testing.T) {
	repo := repositories.NewMemoryRoleRepository()
	svc := NewRoleService(repo)

	role, err := svc.CreateRole(dto.CreateRoleDto{Name: "user"})
	if err != nil {
		t.Fatalf("create error: %v", err)
	}
	found, err := svc.FindRoleByID(role.ID)
	if err != nil || found == nil {
		t.Fatalf("find error: %v", err)
	}
	if found.Name != "user" {
		t.Fatalf("unexpected role name")
	}
}
