package model

// !!! DO NOT EDIT THIS FILE

import (
	"context"
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/mylxsw/eloquent/query"
	"gopkg.in/guregu/null.v3"
	"time"
)

func init() {

}

// StorageFileN is a StorageFile object, all fields are nullable
type StorageFileN struct {
	original         *storageFileOriginal
	storageFileModel *StorageFileModel

	Id        null.Int    `json:"id"`
	Name      null.String `json:"name,omitempty"`
	UserId    null.Int    `json:"user_id,omitempty"`
	FileKey   null.String `json:"file_key,omitempty"`
	Hash      null.String `json:"hash,omitempty"`
	FileSize  null.Int    `json:"file_size,omitempty"`
	Bucket    null.String `json:"bucket,omitempty"`
	Status    null.Int    `json:"status,omitempty"`
	Note      null.String `json:"note,omitempty"`
	Channel   null.String `json:"channel,omitempty"`
	CreatedAt null.Time
	UpdatedAt null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *StorageFileN) As(dst interface{}) error {
	return query.Copy(inst, dst)
}

// SetModel set model for StorageFile
func (inst *StorageFileN) SetModel(storageFileModel *StorageFileModel) {
	inst.storageFileModel = storageFileModel
}

// storageFileOriginal is an object which stores original StorageFile from database
type storageFileOriginal struct {
	Id        null.Int
	Name      null.String
	UserId    null.Int
	FileKey   null.String
	Hash      null.String
	FileSize  null.Int
	Bucket    null.String
	Status    null.Int
	Note      null.String
	Channel   null.String
	CreatedAt null.Time
	UpdatedAt null.Time
}

// Staled identify whether the object has been modified
func (inst *StorageFileN) Staled(onlyFields ...string) bool {
	if inst.original == nil {
		inst.original = &storageFileOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			return true
		}
		if inst.Name != inst.original.Name {
			return true
		}
		if inst.UserId != inst.original.UserId {
			return true
		}
		if inst.FileKey != inst.original.FileKey {
			return true
		}
		if inst.Hash != inst.original.Hash {
			return true
		}
		if inst.FileSize != inst.original.FileSize {
			return true
		}
		if inst.Bucket != inst.original.Bucket {
			return true
		}
		if inst.Status != inst.original.Status {
			return true
		}
		if inst.Note != inst.original.Note {
			return true
		}
		if inst.Channel != inst.original.Channel {
			return true
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			return true
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			return true
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					return true
				}
			case "name":
				if inst.Name != inst.original.Name {
					return true
				}
			case "user_id":
				if inst.UserId != inst.original.UserId {
					return true
				}
			case "file_key":
				if inst.FileKey != inst.original.FileKey {
					return true
				}
			case "hash":
				if inst.Hash != inst.original.Hash {
					return true
				}
			case "file_size":
				if inst.FileSize != inst.original.FileSize {
					return true
				}
			case "bucket":
				if inst.Bucket != inst.original.Bucket {
					return true
				}
			case "status":
				if inst.Status != inst.original.Status {
					return true
				}
			case "note":
				if inst.Note != inst.original.Note {
					return true
				}
			case "channel":
				if inst.Channel != inst.original.Channel {
					return true
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					return true
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					return true
				}
			default:
			}
		}
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *StorageFileN) StaledKV(onlyFields ...string) query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &storageFileOriginal{}
	}

	if len(onlyFields) == 0 {

		if inst.Id != inst.original.Id {
			kv["id"] = inst.Id
		}
		if inst.Name != inst.original.Name {
			kv["name"] = inst.Name
		}
		if inst.UserId != inst.original.UserId {
			kv["user_id"] = inst.UserId
		}
		if inst.FileKey != inst.original.FileKey {
			kv["file_key"] = inst.FileKey
		}
		if inst.Hash != inst.original.Hash {
			kv["hash"] = inst.Hash
		}
		if inst.FileSize != inst.original.FileSize {
			kv["file_size"] = inst.FileSize
		}
		if inst.Bucket != inst.original.Bucket {
			kv["bucket"] = inst.Bucket
		}
		if inst.Status != inst.original.Status {
			kv["status"] = inst.Status
		}
		if inst.Note != inst.original.Note {
			kv["note"] = inst.Note
		}
		if inst.Channel != inst.original.Channel {
			kv["channel"] = inst.Channel
		}
		if inst.CreatedAt != inst.original.CreatedAt {
			kv["created_at"] = inst.CreatedAt
		}
		if inst.UpdatedAt != inst.original.UpdatedAt {
			kv["updated_at"] = inst.UpdatedAt
		}
	} else {
		for _, f := range onlyFields {
			switch strcase.ToSnake(f) {

			case "id":
				if inst.Id != inst.original.Id {
					kv["id"] = inst.Id
				}
			case "name":
				if inst.Name != inst.original.Name {
					kv["name"] = inst.Name
				}
			case "user_id":
				if inst.UserId != inst.original.UserId {
					kv["user_id"] = inst.UserId
				}
			case "file_key":
				if inst.FileKey != inst.original.FileKey {
					kv["file_key"] = inst.FileKey
				}
			case "hash":
				if inst.Hash != inst.original.Hash {
					kv["hash"] = inst.Hash
				}
			case "file_size":
				if inst.FileSize != inst.original.FileSize {
					kv["file_size"] = inst.FileSize
				}
			case "bucket":
				if inst.Bucket != inst.original.Bucket {
					kv["bucket"] = inst.Bucket
				}
			case "status":
				if inst.Status != inst.original.Status {
					kv["status"] = inst.Status
				}
			case "note":
				if inst.Note != inst.original.Note {
					kv["note"] = inst.Note
				}
			case "channel":
				if inst.Channel != inst.original.Channel {
					kv["channel"] = inst.Channel
				}
			case "created_at":
				if inst.CreatedAt != inst.original.CreatedAt {
					kv["created_at"] = inst.CreatedAt
				}
			case "updated_at":
				if inst.UpdatedAt != inst.original.UpdatedAt {
					kv["updated_at"] = inst.UpdatedAt
				}
			default:
			}
		}
	}

	return kv
}

// Save create a new model or update it
func (inst *StorageFileN) Save(ctx context.Context, onlyFields ...string) error {
	if inst.storageFileModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.storageFileModel.SaveOrUpdate(ctx, *inst, onlyFields...)
	if err != nil {
		return err
	}

	inst.Id = null.IntFrom(id)
	return nil
}

// Delete remove a storage_file
func (inst *StorageFileN) Delete(ctx context.Context) error {
	if inst.storageFileModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.storageFileModel.DeleteById(ctx, inst.Id.Int64)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *StorageFileN) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

type storageFileScope struct {
	name  string
	apply func(builder query.Condition)
}

var storageFileGlobalScopes = make([]storageFileScope, 0)
var storageFileLocalScopes = make([]storageFileScope, 0)

// AddGlobalScopeForStorageFile assign a global scope to a model
func AddGlobalScopeForStorageFile(name string, apply func(builder query.Condition)) {
	storageFileGlobalScopes = append(storageFileGlobalScopes, storageFileScope{name: name, apply: apply})
}

// AddLocalScopeForStorageFile assign a local scope to a model
func AddLocalScopeForStorageFile(name string, apply func(builder query.Condition)) {
	storageFileLocalScopes = append(storageFileLocalScopes, storageFileScope{name: name, apply: apply})
}

func (m *StorageFileModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range storageFileGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range storageFileLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *StorageFileModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *StorageFileModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type StorageFile struct {
	Id        int64  `json:"id"`
	Name      string `json:"name,omitempty"`
	UserId    int64  `json:"user_id,omitempty"`
	FileKey   string `json:"file_key,omitempty"`
	Hash      string `json:"hash,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`
	Bucket    string `json:"bucket,omitempty"`
	Status    int64  `json:"status,omitempty"`
	Note      string `json:"note,omitempty"`
	Channel   string `json:"channel,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (w StorageFile) ToStorageFileN(allows ...string) StorageFileN {
	if len(allows) == 0 {
		return StorageFileN{

			Id:        null.IntFrom(int64(w.Id)),
			Name:      null.StringFrom(w.Name),
			UserId:    null.IntFrom(int64(w.UserId)),
			FileKey:   null.StringFrom(w.FileKey),
			Hash:      null.StringFrom(w.Hash),
			FileSize:  null.IntFrom(int64(w.FileSize)),
			Bucket:    null.StringFrom(w.Bucket),
			Status:    null.IntFrom(int64(w.Status)),
			Note:      null.StringFrom(w.Note),
			Channel:   null.StringFrom(w.Channel),
			CreatedAt: null.TimeFrom(w.CreatedAt),
			UpdatedAt: null.TimeFrom(w.UpdatedAt),
		}
	}

	res := StorageFileN{}
	for _, al := range allows {
		switch strcase.ToSnake(al) {

		case "id":
			res.Id = null.IntFrom(int64(w.Id))
		case "name":
			res.Name = null.StringFrom(w.Name)
		case "user_id":
			res.UserId = null.IntFrom(int64(w.UserId))
		case "file_key":
			res.FileKey = null.StringFrom(w.FileKey)
		case "hash":
			res.Hash = null.StringFrom(w.Hash)
		case "file_size":
			res.FileSize = null.IntFrom(int64(w.FileSize))
		case "bucket":
			res.Bucket = null.StringFrom(w.Bucket)
		case "status":
			res.Status = null.IntFrom(int64(w.Status))
		case "note":
			res.Note = null.StringFrom(w.Note)
		case "channel":
			res.Channel = null.StringFrom(w.Channel)
		case "created_at":
			res.CreatedAt = null.TimeFrom(w.CreatedAt)
		case "updated_at":
			res.UpdatedAt = null.TimeFrom(w.UpdatedAt)
		default:
		}
	}

	return res
}

// As convert object to other type
// dst must be a pointer to struct
func (w StorageFile) As(dst interface{}) error {
	return query.Copy(w, dst)
}

func (w *StorageFileN) ToStorageFile() StorageFile {
	return StorageFile{

		Id:        w.Id.Int64,
		Name:      w.Name.String,
		UserId:    w.UserId.Int64,
		FileKey:   w.FileKey.String,
		Hash:      w.Hash.String,
		FileSize:  w.FileSize.Int64,
		Bucket:    w.Bucket.String,
		Status:    w.Status.Int64,
		Note:      w.Note.String,
		Channel:   w.Channel.String,
		CreatedAt: w.CreatedAt.Time,
		UpdatedAt: w.UpdatedAt.Time,
	}
}

// StorageFileModel is a model which encapsulates the operations of the object
type StorageFileModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var storageFileTableName = "storage_file"

// StorageFileTable return table name for StorageFile
func StorageFileTable() string {
	return storageFileTableName
}

const (
	FieldStorageFileId        = "id"
	FieldStorageFileName      = "name"
	FieldStorageFileUserId    = "user_id"
	FieldStorageFileFileKey   = "file_key"
	FieldStorageFileHash      = "hash"
	FieldStorageFileFileSize  = "file_size"
	FieldStorageFileBucket    = "bucket"
	FieldStorageFileStatus    = "status"
	FieldStorageFileNote      = "note"
	FieldStorageFileChannel   = "channel"
	FieldStorageFileCreatedAt = "created_at"
	FieldStorageFileUpdatedAt = "updated_at"
)

// StorageFileFields return all fields in StorageFile model
func StorageFileFields() []string {
	return []string{
		"id",
		"name",
		"user_id",
		"file_key",
		"hash",
		"file_size",
		"bucket",
		"status",
		"note",
		"channel",
		"created_at",
		"updated_at",
	}
}

func SetStorageFileTable(tableName string) {
	storageFileTableName = tableName
}

// NewStorageFileModel create a StorageFileModel
func NewStorageFileModel(db query.Database) *StorageFileModel {
	return &StorageFileModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           storageFileTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *StorageFileModel) GetDB() query.Database {
	return m.db.GetDB()
}

func (m *StorageFileModel) clone() *StorageFileModel {
	return &StorageFileModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *StorageFileModel) WithoutGlobalScopes(names ...string) *StorageFileModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *StorageFileModel) WithLocalScopes(names ...string) *StorageFileModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *StorageFileModel) Condition(builder query.SQLBuilder) *StorageFileModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *StorageFileModel) Find(ctx context.Context, id int64) (*StorageFileN, error) {
	return m.First(ctx, m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *StorageFileModel) Exists(ctx context.Context, builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(ctx, builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *StorageFileModel) Count(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {
	sqlStr, params := m.query.
		Merge(builders...).
		Table(m.tableName).
		AppendCondition(m.applyScope()).
		ResolveCount()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()
	var res int64
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (m *StorageFileModel) Paginate(ctx context.Context, page int64, perPage int64, builders ...query.SQLBuilder) ([]StorageFileN, query.PaginateMeta, error) {
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 15
	}

	meta := query.PaginateMeta{
		PerPage: perPage,
		Page:    page,
	}

	count, err := m.Count(ctx, builders...)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = count
	meta.LastPage = count / perPage
	if count%perPage != 0 {
		meta.LastPage += 1
	}

	res, err := m.Get(ctx, append([]query.SQLBuilder{query.Builder().Limit(perPage).Offset((page - 1) * perPage)}, builders...)...)
	if err != nil {
		return res, meta, err
	}

	return res, meta, nil
}

// Get retrieve all results for given query
func (m *StorageFileModel) Get(ctx context.Context, builders ...query.SQLBuilder) ([]StorageFileN, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"name",
			"user_id",
			"file_key",
			"hash",
			"file_size",
			"bucket",
			"status",
			"note",
			"channel",
			"created_at",
			"updated_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "id":
			selectFields = append(selectFields, f)
		case "name":
			selectFields = append(selectFields, f)
		case "user_id":
			selectFields = append(selectFields, f)
		case "file_key":
			selectFields = append(selectFields, f)
		case "hash":
			selectFields = append(selectFields, f)
		case "file_size":
			selectFields = append(selectFields, f)
		case "bucket":
			selectFields = append(selectFields, f)
		case "status":
			selectFields = append(selectFields, f)
		case "note":
			selectFields = append(selectFields, f)
		case "channel":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*StorageFileN, []interface{}) {
		var storageFileVar StorageFileN
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &storageFileVar.Id)
			case "name":
				scanFields = append(scanFields, &storageFileVar.Name)
			case "user_id":
				scanFields = append(scanFields, &storageFileVar.UserId)
			case "file_key":
				scanFields = append(scanFields, &storageFileVar.FileKey)
			case "hash":
				scanFields = append(scanFields, &storageFileVar.Hash)
			case "file_size":
				scanFields = append(scanFields, &storageFileVar.FileSize)
			case "bucket":
				scanFields = append(scanFields, &storageFileVar.Bucket)
			case "status":
				scanFields = append(scanFields, &storageFileVar.Status)
			case "note":
				scanFields = append(scanFields, &storageFileVar.Note)
			case "channel":
				scanFields = append(scanFields, &storageFileVar.Channel)
			case "created_at":
				scanFields = append(scanFields, &storageFileVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &storageFileVar.UpdatedAt)
			}
		}

		return &storageFileVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(ctx, sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	storageFiles := make([]StorageFileN, 0)
	for rows.Next() {
		storageFileReal, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		storageFileReal.original = &storageFileOriginal{}
		_ = query.Copy(storageFileReal, storageFileReal.original)

		storageFileReal.SetModel(m)
		storageFiles = append(storageFiles, *storageFileReal)
	}

	return storageFiles, nil
}

// First return first result for given query
func (m *StorageFileModel) First(ctx context.Context, builders ...query.SQLBuilder) (*StorageFileN, error) {
	res, err := m.Get(ctx, append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, query.ErrNoResult
	}

	return &res[0], nil
}

// Create save a new storage_file to database
func (m *StorageFileModel) Create(ctx context.Context, kv query.KV) (int64, error) {

	if _, ok := kv["created_at"]; !ok {
		kv["created_at"] = time.Now()
	}

	if _, ok := kv["updated_at"]; !ok {
		kv["updated_at"] = time.Now()
	}

	sqlStr, params := m.query.Table(m.tableName).ResolveInsert(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// SaveAll save all storage_files to database
func (m *StorageFileModel) SaveAll(ctx context.Context, storageFiles []StorageFileN) ([]int64, error) {
	ids := make([]int64, 0)
	for _, storageFile := range storageFiles {
		id, err := m.Save(ctx, storageFile)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a storage_file to database
func (m *StorageFileModel) Save(ctx context.Context, storageFile StorageFileN, onlyFields ...string) (int64, error) {
	return m.Create(ctx, storageFile.StaledKV(onlyFields...))
}

// SaveOrUpdate save a new storage_file or update it when it has a id > 0
func (m *StorageFileModel) SaveOrUpdate(ctx context.Context, storageFile StorageFileN, onlyFields ...string) (id int64, updated bool, err error) {
	if storageFile.Id.Int64 > 0 {
		_, _err := m.UpdateById(ctx, storageFile.Id.Int64, storageFile, onlyFields...)
		return storageFile.Id.Int64, true, _err
	}

	_id, _err := m.Save(ctx, storageFile, onlyFields...)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *StorageFileModel) UpdateFields(ctx context.Context, kv query.KV, builders ...query.SQLBuilder) (int64, error) {
	if len(kv) == 0 {
		return 0, nil
	}

	kv["updated_at"] = time.Now()

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).
		Table(m.tableName).
		ResolveUpdate(kv)

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Update update a model for given query
func (m *StorageFileModel) Update(ctx context.Context, builder query.SQLBuilder, storageFile StorageFileN, onlyFields ...string) (int64, error) {
	return m.UpdateFields(ctx, storageFile.StaledKV(onlyFields...), builder)
}

// UpdateById update a model by id
func (m *StorageFileModel) UpdateById(ctx context.Context, id int64, storageFile StorageFileN, onlyFields ...string) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).UpdateFields(ctx, storageFile.StaledKV(onlyFields...))
}

// Delete remove a model
func (m *StorageFileModel) Delete(ctx context.Context, builders ...query.SQLBuilder) (int64, error) {

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).Table(m.tableName).ResolveDelete()

	res, err := m.db.ExecContext(ctx, sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

// DeleteById remove a model by id
func (m *StorageFileModel) DeleteById(ctx context.Context, id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete(ctx)
}
