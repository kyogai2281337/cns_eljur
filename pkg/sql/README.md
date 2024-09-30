***БД***

Structure:

```go
type Cabinet struct {
	ID       int64   `bson:"_id,omitempty" json:"id"`
	Name     string  `json:"name"`
	Type     CabType `json:"type"`
	Capacity int     `json:"capacity"`
}

type Group struct {
	ID             int64           `bson:"_id,omitempty" json:"id"`
	Specialization *Specialization `json:"specialization"`
	Name           string          `json:"name"`
	MaxPairs       int             `json:"max_pairs"`
}

type Role struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Specialization struct {
	ID        int64            `bson:"_id,omitempty" json:"id"`
	Name      string           `json:"name"`
	Course    int              `json:"course"`
	EduPlan   map[*Subject]int `json:"-"`
	ShortPlan map[int64]int    `json:"short_plan"`
	PlanId    string           `json:"plan_id"`
}

type Subject struct {
	ID               int64   `bson:"_id,omitempty" json:"id"`
	Name             string  `json:"name"`
	RecommendCabType CabType `json:"type"`
}

type Teacher struct {
	ID               int64                 `json:"id"`
	Name             string                `json:"name"`
	Links            map[*Group][]*Subject `json:"full_links"`
	LinksID          string                `bson:"_id" json:"links_id"`
	RecommendSchCap_ int                   `json:"capacity"`
	SL               map[int64][]int64     `bson:"links" json:"links"`
}

type User struct {
	ID        int64  `json:"id,omitempty"`
	Email     string `json:"email"`
	Pass      string `json:"password,omitempty"`
	EncPass   string `json:"-"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	IsActive  bool   `json:"-"`
	Role      *Role  `json:"role,omitempty"`
}
```

На некоторых из этих структур возложен экстра-ответственность, в виде валидаций или превращений полей из длинного в
короткий тип.

С методологической точки зрения всё просто, реализованы базовые CRUD-операции:

```go
type AnyRepository interface {
	Find(int64) (*model.Any, error)
	Create(context.Context, *model.Any) (*model.Any, error)
	FindByName(string) (*model.Any, error)
	GetList(page int64, limit int64) ([]*model.Any, error)
	Update(context.Context, *model.Any) error
}
```

Так же реализована механика транзакций, для увеличения масштабируемости и предотвращения проблем с отменами операций:

```go
func (s *Store) BeginTx(ctx context.Context) (context.Context, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("database store error:%s", err.Error())
	}
	return context.WithValue(ctx, txKey{}, tx), nil
}

func (s *Store) CommitTx(ctx context.Context) error {
	tx, ok := ctx.Value(txKey{}).(*sql.Tx)
	if !ok {
		return errors.New("cannot commit: no transaction in context") // TODO: static err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("database store error:%s", err.Error())
	}
	return nil
}

func (s *Store) RollbackTx(ctx context.Context) error {
	tx, ok := ctx.Value(txKey{}).(*sql.Tx)
	if !ok {
		return errors.New("cannot rollback: no transaction in context") // TODO: static err
	}
	if err := tx.Rollback(); err != nil {
		return fmt.Errorf("database store error:%s", err.Error())
	}
	return nil
}

func (s *Store) getTxFromCtx(ctx context.Context) (*sql.Tx, error) {
	tx, ok := ctx.Value(txKey{}).(*sql.Tx)
	if !ok {
		return nil, errors.New("no transaction in context")
	}
	return tx, nil
}

```

Подключение реализовано через паттерн **Стратегия**, что позволяет легко и просто, подготовить и перенести старую БД на
новую, или реализовать тестовое хранилище in-memory. Необходимо всего лишь реализовать интерфейс текущей реализации в
новой БД.
