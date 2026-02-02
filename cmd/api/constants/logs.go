package constants

// Layer
const (
	LayerController = "controller"
	LayerGateway    = "gateway"
	LayerRepository = "repository"
	LayerService    = "service"
)

// Module
const (
	ModuleCategories          = "categories "
	ModuleSkills              = "skills"
	ModuleSkillsAndCategories = "skills_categories"
	ModuleTasker              = "tasker"
)

// Function Skills
const (
	FunctionSkillsFindBy  = "skills_find_by"
	FunctionSkillsFindAll = "skills_find_all"
	FunctionSkillsFirstBy = "skills_first_by"
	FunctionSkillsSearch  = "skills_search"
	FunctionSkillsList    = "skills_list"
	FunctionSkillsSave    = "skills_save"
	FunctionSkillsUpdate  = "skills_update"
	FunctionSkillsUpsert  = "skills_upsert"
)

// Function Categories
const (
	FunctionCategorySearch    = "category_search"
	FunctionCategoryList      = "category_list"
	FunctionCategoriesFindAll = "category_find_all"
	FunctionCategoriesFirstBy = "category_first_by"
	FunctionCategoriesUpsert  = "category_upsert"
	FunctionCategoriesSave    = "category_save"
)

// Function Skills-Categories
const (
	FunctionSkillsCategoriesFirstBy = "skills_category_first_by"
	FunctionSkillsCategoriesSave    = "skills_category_save"
	FunctionSkillsCategoriesUpdate  = "skills_category_update"
	FunctionSkillsCategoriesUpsert  = "skills_category_upsert"
)

// Function Tasker
const (
	FunctionTaskerList = "tasker_list"
)
