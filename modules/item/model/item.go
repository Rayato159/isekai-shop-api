package model

type (
	Item struct {
		ID          uint64  `json:"id"`
		AdminID     *string `json:"adminID"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Picture     string  `json:"picture"`
		Price       uint    `json:"price"`
	}

	ItemFilter struct {
		Name        string `query:"name" validate:"omitempty,max=64"`
		Description string `query:"description" validate:"omitempty,max=128"`
		Paginate
	}

	Paginate struct {
		Page int64 `query:"page" validate:"required,min=1"`
		Size int64 `query:"size" validate:"required,min=1 max=20"`
	}

	ItemResult struct {
		Items    []*Item        `json:"items"`
		Paginate PaginateResult `json:"paginate"`
	}

	PaginateResult struct {
		Page      int64 `json:"page"`
		TotalPage int64 `json:"totalPage"`
	}

	CreateItemReq struct {
		AdminID     string
		Name        string `json:"name" validate:"required,max=64"`
		Description string `json:"description" validate:"required,max=128"`
		Picture     string `json:"picture" validate:"required"`
		Price       uint   `json:"price" validate:"required"`
	}

	EditItemReq struct {
		AdminID     string
		Name        string `json:"name" validate:"omitempty,max=64"`
		Description string `json:"description" validate:"omitempty,max=128"`
		Picture     string `json:"picture" validate:"omitempty"`
		Price       uint   `json:"price" validate:"omitempty"`
	}
)
