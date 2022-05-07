// Code generated from TSqlParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // TSqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTSqlParserListener is a complete listener for a parse tree produced by TSqlParser.
type BaseTSqlParserListener struct{}

var _ TSqlParserListener = &BaseTSqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTSqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTSqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTSqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTSqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTsql_file is called when production tsql_file is entered.
func (s *BaseTSqlParserListener) EnterTsql_file(ctx *Tsql_fileContext) {}

// ExitTsql_file is called when production tsql_file is exited.
func (s *BaseTSqlParserListener) ExitTsql_file(ctx *Tsql_fileContext) {}

// EnterBatch is called when production batch is entered.
func (s *BaseTSqlParserListener) EnterBatch(ctx *BatchContext) {}

// ExitBatch is called when production batch is exited.
func (s *BaseTSqlParserListener) ExitBatch(ctx *BatchContext) {}

// EnterBatch_level_statement is called when production batch_level_statement is entered.
func (s *BaseTSqlParserListener) EnterBatch_level_statement(ctx *Batch_level_statementContext) {}

// ExitBatch_level_statement is called when production batch_level_statement is exited.
func (s *BaseTSqlParserListener) ExitBatch_level_statement(ctx *Batch_level_statementContext) {}

// EnterSql_clauses is called when production sql_clauses is entered.
func (s *BaseTSqlParserListener) EnterSql_clauses(ctx *Sql_clausesContext) {}

// ExitSql_clauses is called when production sql_clauses is exited.
func (s *BaseTSqlParserListener) ExitSql_clauses(ctx *Sql_clausesContext) {}

// EnterDml_clause is called when production dml_clause is entered.
func (s *BaseTSqlParserListener) EnterDml_clause(ctx *Dml_clauseContext) {}

// ExitDml_clause is called when production dml_clause is exited.
func (s *BaseTSqlParserListener) ExitDml_clause(ctx *Dml_clauseContext) {}

// EnterDdl_clause is called when production ddl_clause is entered.
func (s *BaseTSqlParserListener) EnterDdl_clause(ctx *Ddl_clauseContext) {}

// ExitDdl_clause is called when production ddl_clause is exited.
func (s *BaseTSqlParserListener) ExitDdl_clause(ctx *Ddl_clauseContext) {}

// EnterBackup_statement is called when production backup_statement is entered.
func (s *BaseTSqlParserListener) EnterBackup_statement(ctx *Backup_statementContext) {}

// ExitBackup_statement is called when production backup_statement is exited.
func (s *BaseTSqlParserListener) ExitBackup_statement(ctx *Backup_statementContext) {}

// EnterCfl_statement is called when production cfl_statement is entered.
func (s *BaseTSqlParserListener) EnterCfl_statement(ctx *Cfl_statementContext) {}

// ExitCfl_statement is called when production cfl_statement is exited.
func (s *BaseTSqlParserListener) ExitCfl_statement(ctx *Cfl_statementContext) {}

// EnterBlock_statement is called when production block_statement is entered.
func (s *BaseTSqlParserListener) EnterBlock_statement(ctx *Block_statementContext) {}

// ExitBlock_statement is called when production block_statement is exited.
func (s *BaseTSqlParserListener) ExitBlock_statement(ctx *Block_statementContext) {}

// EnterBreak_statement is called when production break_statement is entered.
func (s *BaseTSqlParserListener) EnterBreak_statement(ctx *Break_statementContext) {}

// ExitBreak_statement is called when production break_statement is exited.
func (s *BaseTSqlParserListener) ExitBreak_statement(ctx *Break_statementContext) {}

// EnterContinue_statement is called when production continue_statement is entered.
func (s *BaseTSqlParserListener) EnterContinue_statement(ctx *Continue_statementContext) {}

// ExitContinue_statement is called when production continue_statement is exited.
func (s *BaseTSqlParserListener) ExitContinue_statement(ctx *Continue_statementContext) {}

// EnterGoto_statement is called when production goto_statement is entered.
func (s *BaseTSqlParserListener) EnterGoto_statement(ctx *Goto_statementContext) {}

// ExitGoto_statement is called when production goto_statement is exited.
func (s *BaseTSqlParserListener) ExitGoto_statement(ctx *Goto_statementContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseTSqlParserListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseTSqlParserListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseTSqlParserListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseTSqlParserListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterThrow_statement is called when production throw_statement is entered.
func (s *BaseTSqlParserListener) EnterThrow_statement(ctx *Throw_statementContext) {}

// ExitThrow_statement is called when production throw_statement is exited.
func (s *BaseTSqlParserListener) ExitThrow_statement(ctx *Throw_statementContext) {}

// EnterThrow_error_number is called when production throw_error_number is entered.
func (s *BaseTSqlParserListener) EnterThrow_error_number(ctx *Throw_error_numberContext) {}

// ExitThrow_error_number is called when production throw_error_number is exited.
func (s *BaseTSqlParserListener) ExitThrow_error_number(ctx *Throw_error_numberContext) {}

// EnterThrow_message is called when production throw_message is entered.
func (s *BaseTSqlParserListener) EnterThrow_message(ctx *Throw_messageContext) {}

// ExitThrow_message is called when production throw_message is exited.
func (s *BaseTSqlParserListener) ExitThrow_message(ctx *Throw_messageContext) {}

// EnterThrow_state is called when production throw_state is entered.
func (s *BaseTSqlParserListener) EnterThrow_state(ctx *Throw_stateContext) {}

// ExitThrow_state is called when production throw_state is exited.
func (s *BaseTSqlParserListener) ExitThrow_state(ctx *Throw_stateContext) {}

// EnterTry_catch_statement is called when production try_catch_statement is entered.
func (s *BaseTSqlParserListener) EnterTry_catch_statement(ctx *Try_catch_statementContext) {}

// ExitTry_catch_statement is called when production try_catch_statement is exited.
func (s *BaseTSqlParserListener) ExitTry_catch_statement(ctx *Try_catch_statementContext) {}

// EnterWaitfor_statement is called when production waitfor_statement is entered.
func (s *BaseTSqlParserListener) EnterWaitfor_statement(ctx *Waitfor_statementContext) {}

// ExitWaitfor_statement is called when production waitfor_statement is exited.
func (s *BaseTSqlParserListener) ExitWaitfor_statement(ctx *Waitfor_statementContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BaseTSqlParserListener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BaseTSqlParserListener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterPrint_statement is called when production print_statement is entered.
func (s *BaseTSqlParserListener) EnterPrint_statement(ctx *Print_statementContext) {}

// ExitPrint_statement is called when production print_statement is exited.
func (s *BaseTSqlParserListener) ExitPrint_statement(ctx *Print_statementContext) {}

// EnterRaiseerror_statement is called when production raiseerror_statement is entered.
func (s *BaseTSqlParserListener) EnterRaiseerror_statement(ctx *Raiseerror_statementContext) {}

// ExitRaiseerror_statement is called when production raiseerror_statement is exited.
func (s *BaseTSqlParserListener) ExitRaiseerror_statement(ctx *Raiseerror_statementContext) {}

// EnterEmpty_statement is called when production empty_statement is entered.
func (s *BaseTSqlParserListener) EnterEmpty_statement(ctx *Empty_statementContext) {}

// ExitEmpty_statement is called when production empty_statement is exited.
func (s *BaseTSqlParserListener) ExitEmpty_statement(ctx *Empty_statementContext) {}

// EnterAnother_statement is called when production another_statement is entered.
func (s *BaseTSqlParserListener) EnterAnother_statement(ctx *Another_statementContext) {}

// ExitAnother_statement is called when production another_statement is exited.
func (s *BaseTSqlParserListener) ExitAnother_statement(ctx *Another_statementContext) {}

// EnterAlter_application_role is called when production alter_application_role is entered.
func (s *BaseTSqlParserListener) EnterAlter_application_role(ctx *Alter_application_roleContext) {}

// ExitAlter_application_role is called when production alter_application_role is exited.
func (s *BaseTSqlParserListener) ExitAlter_application_role(ctx *Alter_application_roleContext) {}

// EnterCreate_application_role is called when production create_application_role is entered.
func (s *BaseTSqlParserListener) EnterCreate_application_role(ctx *Create_application_roleContext) {}

// ExitCreate_application_role is called when production create_application_role is exited.
func (s *BaseTSqlParserListener) ExitCreate_application_role(ctx *Create_application_roleContext) {}

// EnterDrop_aggregate is called when production drop_aggregate is entered.
func (s *BaseTSqlParserListener) EnterDrop_aggregate(ctx *Drop_aggregateContext) {}

// ExitDrop_aggregate is called when production drop_aggregate is exited.
func (s *BaseTSqlParserListener) ExitDrop_aggregate(ctx *Drop_aggregateContext) {}

// EnterDrop_application_role is called when production drop_application_role is entered.
func (s *BaseTSqlParserListener) EnterDrop_application_role(ctx *Drop_application_roleContext) {}

// ExitDrop_application_role is called when production drop_application_role is exited.
func (s *BaseTSqlParserListener) ExitDrop_application_role(ctx *Drop_application_roleContext) {}

// EnterAlter_assembly is called when production alter_assembly is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly(ctx *Alter_assemblyContext) {}

// ExitAlter_assembly is called when production alter_assembly is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly(ctx *Alter_assemblyContext) {}

// EnterAlter_assembly_start is called when production alter_assembly_start is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_start(ctx *Alter_assembly_startContext) {}

// ExitAlter_assembly_start is called when production alter_assembly_start is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_start(ctx *Alter_assembly_startContext) {}

// EnterAlter_assembly_clause is called when production alter_assembly_clause is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_clause(ctx *Alter_assembly_clauseContext) {}

// ExitAlter_assembly_clause is called when production alter_assembly_clause is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_clause(ctx *Alter_assembly_clauseContext) {}

// EnterAlter_assembly_from_clause is called when production alter_assembly_from_clause is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_from_clause(ctx *Alter_assembly_from_clauseContext) {
}

// ExitAlter_assembly_from_clause is called when production alter_assembly_from_clause is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_from_clause(ctx *Alter_assembly_from_clauseContext) {
}

// EnterAlter_assembly_from_clause_start is called when production alter_assembly_from_clause_start is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_from_clause_start(ctx *Alter_assembly_from_clause_startContext) {
}

// ExitAlter_assembly_from_clause_start is called when production alter_assembly_from_clause_start is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_from_clause_start(ctx *Alter_assembly_from_clause_startContext) {
}

// EnterAlter_assembly_drop_clause is called when production alter_assembly_drop_clause is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_drop_clause(ctx *Alter_assembly_drop_clauseContext) {
}

// ExitAlter_assembly_drop_clause is called when production alter_assembly_drop_clause is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_drop_clause(ctx *Alter_assembly_drop_clauseContext) {
}

// EnterAlter_assembly_drop_multiple_files is called when production alter_assembly_drop_multiple_files is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_drop_multiple_files(ctx *Alter_assembly_drop_multiple_filesContext) {
}

// ExitAlter_assembly_drop_multiple_files is called when production alter_assembly_drop_multiple_files is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_drop_multiple_files(ctx *Alter_assembly_drop_multiple_filesContext) {
}

// EnterAlter_assembly_drop is called when production alter_assembly_drop is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_drop(ctx *Alter_assembly_dropContext) {}

// ExitAlter_assembly_drop is called when production alter_assembly_drop is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_drop(ctx *Alter_assembly_dropContext) {}

// EnterAlter_assembly_add_clause is called when production alter_assembly_add_clause is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_add_clause(ctx *Alter_assembly_add_clauseContext) {
}

// ExitAlter_assembly_add_clause is called when production alter_assembly_add_clause is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_add_clause(ctx *Alter_assembly_add_clauseContext) {
}

// EnterAlter_asssembly_add_clause_start is called when production alter_asssembly_add_clause_start is entered.
func (s *BaseTSqlParserListener) EnterAlter_asssembly_add_clause_start(ctx *Alter_asssembly_add_clause_startContext) {
}

// ExitAlter_asssembly_add_clause_start is called when production alter_asssembly_add_clause_start is exited.
func (s *BaseTSqlParserListener) ExitAlter_asssembly_add_clause_start(ctx *Alter_asssembly_add_clause_startContext) {
}

// EnterAlter_assembly_client_file_clause is called when production alter_assembly_client_file_clause is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_client_file_clause(ctx *Alter_assembly_client_file_clauseContext) {
}

// ExitAlter_assembly_client_file_clause is called when production alter_assembly_client_file_clause is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_client_file_clause(ctx *Alter_assembly_client_file_clauseContext) {
}

// EnterAlter_assembly_file_name is called when production alter_assembly_file_name is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_file_name(ctx *Alter_assembly_file_nameContext) {
}

// ExitAlter_assembly_file_name is called when production alter_assembly_file_name is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_file_name(ctx *Alter_assembly_file_nameContext) {}

// EnterAlter_assembly_file_bits is called when production alter_assembly_file_bits is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_file_bits(ctx *Alter_assembly_file_bitsContext) {
}

// ExitAlter_assembly_file_bits is called when production alter_assembly_file_bits is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_file_bits(ctx *Alter_assembly_file_bitsContext) {}

// EnterAlter_assembly_as is called when production alter_assembly_as is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_as(ctx *Alter_assembly_asContext) {}

// ExitAlter_assembly_as is called when production alter_assembly_as is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_as(ctx *Alter_assembly_asContext) {}

// EnterAlter_assembly_with_clause is called when production alter_assembly_with_clause is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_with_clause(ctx *Alter_assembly_with_clauseContext) {
}

// ExitAlter_assembly_with_clause is called when production alter_assembly_with_clause is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_with_clause(ctx *Alter_assembly_with_clauseContext) {
}

// EnterAlter_assembly_with is called when production alter_assembly_with is entered.
func (s *BaseTSqlParserListener) EnterAlter_assembly_with(ctx *Alter_assembly_withContext) {}

// ExitAlter_assembly_with is called when production alter_assembly_with is exited.
func (s *BaseTSqlParserListener) ExitAlter_assembly_with(ctx *Alter_assembly_withContext) {}

// EnterClient_assembly_specifier is called when production client_assembly_specifier is entered.
func (s *BaseTSqlParserListener) EnterClient_assembly_specifier(ctx *Client_assembly_specifierContext) {
}

// ExitClient_assembly_specifier is called when production client_assembly_specifier is exited.
func (s *BaseTSqlParserListener) ExitClient_assembly_specifier(ctx *Client_assembly_specifierContext) {
}

// EnterAssembly_option is called when production assembly_option is entered.
func (s *BaseTSqlParserListener) EnterAssembly_option(ctx *Assembly_optionContext) {}

// ExitAssembly_option is called when production assembly_option is exited.
func (s *BaseTSqlParserListener) ExitAssembly_option(ctx *Assembly_optionContext) {}

// EnterNetwork_file_share is called when production network_file_share is entered.
func (s *BaseTSqlParserListener) EnterNetwork_file_share(ctx *Network_file_shareContext) {}

// ExitNetwork_file_share is called when production network_file_share is exited.
func (s *BaseTSqlParserListener) ExitNetwork_file_share(ctx *Network_file_shareContext) {}

// EnterNetwork_computer is called when production network_computer is entered.
func (s *BaseTSqlParserListener) EnterNetwork_computer(ctx *Network_computerContext) {}

// ExitNetwork_computer is called when production network_computer is exited.
func (s *BaseTSqlParserListener) ExitNetwork_computer(ctx *Network_computerContext) {}

// EnterNetwork_file_start is called when production network_file_start is entered.
func (s *BaseTSqlParserListener) EnterNetwork_file_start(ctx *Network_file_startContext) {}

// ExitNetwork_file_start is called when production network_file_start is exited.
func (s *BaseTSqlParserListener) ExitNetwork_file_start(ctx *Network_file_startContext) {}

// EnterFile_path is called when production file_path is entered.
func (s *BaseTSqlParserListener) EnterFile_path(ctx *File_pathContext) {}

// ExitFile_path is called when production file_path is exited.
func (s *BaseTSqlParserListener) ExitFile_path(ctx *File_pathContext) {}

// EnterFile_directory_path_separator is called when production file_directory_path_separator is entered.
func (s *BaseTSqlParserListener) EnterFile_directory_path_separator(ctx *File_directory_path_separatorContext) {
}

// ExitFile_directory_path_separator is called when production file_directory_path_separator is exited.
func (s *BaseTSqlParserListener) ExitFile_directory_path_separator(ctx *File_directory_path_separatorContext) {
}

// EnterLocal_file is called when production local_file is entered.
func (s *BaseTSqlParserListener) EnterLocal_file(ctx *Local_fileContext) {}

// ExitLocal_file is called when production local_file is exited.
func (s *BaseTSqlParserListener) ExitLocal_file(ctx *Local_fileContext) {}

// EnterLocal_drive is called when production local_drive is entered.
func (s *BaseTSqlParserListener) EnterLocal_drive(ctx *Local_driveContext) {}

// ExitLocal_drive is called when production local_drive is exited.
func (s *BaseTSqlParserListener) ExitLocal_drive(ctx *Local_driveContext) {}

// EnterMultiple_local_files is called when production multiple_local_files is entered.
func (s *BaseTSqlParserListener) EnterMultiple_local_files(ctx *Multiple_local_filesContext) {}

// ExitMultiple_local_files is called when production multiple_local_files is exited.
func (s *BaseTSqlParserListener) ExitMultiple_local_files(ctx *Multiple_local_filesContext) {}

// EnterMultiple_local_file_start is called when production multiple_local_file_start is entered.
func (s *BaseTSqlParserListener) EnterMultiple_local_file_start(ctx *Multiple_local_file_startContext) {
}

// ExitMultiple_local_file_start is called when production multiple_local_file_start is exited.
func (s *BaseTSqlParserListener) ExitMultiple_local_file_start(ctx *Multiple_local_file_startContext) {
}

// EnterCreate_assembly is called when production create_assembly is entered.
func (s *BaseTSqlParserListener) EnterCreate_assembly(ctx *Create_assemblyContext) {}

// ExitCreate_assembly is called when production create_assembly is exited.
func (s *BaseTSqlParserListener) ExitCreate_assembly(ctx *Create_assemblyContext) {}

// EnterDrop_assembly is called when production drop_assembly is entered.
func (s *BaseTSqlParserListener) EnterDrop_assembly(ctx *Drop_assemblyContext) {}

// ExitDrop_assembly is called when production drop_assembly is exited.
func (s *BaseTSqlParserListener) ExitDrop_assembly(ctx *Drop_assemblyContext) {}

// EnterAlter_asymmetric_key is called when production alter_asymmetric_key is entered.
func (s *BaseTSqlParserListener) EnterAlter_asymmetric_key(ctx *Alter_asymmetric_keyContext) {}

// ExitAlter_asymmetric_key is called when production alter_asymmetric_key is exited.
func (s *BaseTSqlParserListener) ExitAlter_asymmetric_key(ctx *Alter_asymmetric_keyContext) {}

// EnterAlter_asymmetric_key_start is called when production alter_asymmetric_key_start is entered.
func (s *BaseTSqlParserListener) EnterAlter_asymmetric_key_start(ctx *Alter_asymmetric_key_startContext) {
}

// ExitAlter_asymmetric_key_start is called when production alter_asymmetric_key_start is exited.
func (s *BaseTSqlParserListener) ExitAlter_asymmetric_key_start(ctx *Alter_asymmetric_key_startContext) {
}

// EnterAsymmetric_key_option is called when production asymmetric_key_option is entered.
func (s *BaseTSqlParserListener) EnterAsymmetric_key_option(ctx *Asymmetric_key_optionContext) {}

// ExitAsymmetric_key_option is called when production asymmetric_key_option is exited.
func (s *BaseTSqlParserListener) ExitAsymmetric_key_option(ctx *Asymmetric_key_optionContext) {}

// EnterAsymmetric_key_option_start is called when production asymmetric_key_option_start is entered.
func (s *BaseTSqlParserListener) EnterAsymmetric_key_option_start(ctx *Asymmetric_key_option_startContext) {
}

// ExitAsymmetric_key_option_start is called when production asymmetric_key_option_start is exited.
func (s *BaseTSqlParserListener) ExitAsymmetric_key_option_start(ctx *Asymmetric_key_option_startContext) {
}

// EnterAsymmetric_key_password_change_option is called when production asymmetric_key_password_change_option is entered.
func (s *BaseTSqlParserListener) EnterAsymmetric_key_password_change_option(ctx *Asymmetric_key_password_change_optionContext) {
}

// ExitAsymmetric_key_password_change_option is called when production asymmetric_key_password_change_option is exited.
func (s *BaseTSqlParserListener) ExitAsymmetric_key_password_change_option(ctx *Asymmetric_key_password_change_optionContext) {
}

// EnterCreate_asymmetric_key is called when production create_asymmetric_key is entered.
func (s *BaseTSqlParserListener) EnterCreate_asymmetric_key(ctx *Create_asymmetric_keyContext) {}

// ExitCreate_asymmetric_key is called when production create_asymmetric_key is exited.
func (s *BaseTSqlParserListener) ExitCreate_asymmetric_key(ctx *Create_asymmetric_keyContext) {}

// EnterDrop_asymmetric_key is called when production drop_asymmetric_key is entered.
func (s *BaseTSqlParserListener) EnterDrop_asymmetric_key(ctx *Drop_asymmetric_keyContext) {}

// ExitDrop_asymmetric_key is called when production drop_asymmetric_key is exited.
func (s *BaseTSqlParserListener) ExitDrop_asymmetric_key(ctx *Drop_asymmetric_keyContext) {}

// EnterAlter_authorization is called when production alter_authorization is entered.
func (s *BaseTSqlParserListener) EnterAlter_authorization(ctx *Alter_authorizationContext) {}

// ExitAlter_authorization is called when production alter_authorization is exited.
func (s *BaseTSqlParserListener) ExitAlter_authorization(ctx *Alter_authorizationContext) {}

// EnterAuthorization_grantee is called when production authorization_grantee is entered.
func (s *BaseTSqlParserListener) EnterAuthorization_grantee(ctx *Authorization_granteeContext) {}

// ExitAuthorization_grantee is called when production authorization_grantee is exited.
func (s *BaseTSqlParserListener) ExitAuthorization_grantee(ctx *Authorization_granteeContext) {}

// EnterEntity_to is called when production entity_to is entered.
func (s *BaseTSqlParserListener) EnterEntity_to(ctx *Entity_toContext) {}

// ExitEntity_to is called when production entity_to is exited.
func (s *BaseTSqlParserListener) ExitEntity_to(ctx *Entity_toContext) {}

// EnterColon_colon is called when production colon_colon is entered.
func (s *BaseTSqlParserListener) EnterColon_colon(ctx *Colon_colonContext) {}

// ExitColon_colon is called when production colon_colon is exited.
func (s *BaseTSqlParserListener) ExitColon_colon(ctx *Colon_colonContext) {}

// EnterAlter_authorization_start is called when production alter_authorization_start is entered.
func (s *BaseTSqlParserListener) EnterAlter_authorization_start(ctx *Alter_authorization_startContext) {
}

// ExitAlter_authorization_start is called when production alter_authorization_start is exited.
func (s *BaseTSqlParserListener) ExitAlter_authorization_start(ctx *Alter_authorization_startContext) {
}

// EnterAlter_authorization_for_sql_database is called when production alter_authorization_for_sql_database is entered.
func (s *BaseTSqlParserListener) EnterAlter_authorization_for_sql_database(ctx *Alter_authorization_for_sql_databaseContext) {
}

// ExitAlter_authorization_for_sql_database is called when production alter_authorization_for_sql_database is exited.
func (s *BaseTSqlParserListener) ExitAlter_authorization_for_sql_database(ctx *Alter_authorization_for_sql_databaseContext) {
}

// EnterAlter_authorization_for_azure_dw is called when production alter_authorization_for_azure_dw is entered.
func (s *BaseTSqlParserListener) EnterAlter_authorization_for_azure_dw(ctx *Alter_authorization_for_azure_dwContext) {
}

// ExitAlter_authorization_for_azure_dw is called when production alter_authorization_for_azure_dw is exited.
func (s *BaseTSqlParserListener) ExitAlter_authorization_for_azure_dw(ctx *Alter_authorization_for_azure_dwContext) {
}

// EnterAlter_authorization_for_parallel_dw is called when production alter_authorization_for_parallel_dw is entered.
func (s *BaseTSqlParserListener) EnterAlter_authorization_for_parallel_dw(ctx *Alter_authorization_for_parallel_dwContext) {
}

// ExitAlter_authorization_for_parallel_dw is called when production alter_authorization_for_parallel_dw is exited.
func (s *BaseTSqlParserListener) ExitAlter_authorization_for_parallel_dw(ctx *Alter_authorization_for_parallel_dwContext) {
}

// EnterClass_type is called when production class_type is entered.
func (s *BaseTSqlParserListener) EnterClass_type(ctx *Class_typeContext) {}

// ExitClass_type is called when production class_type is exited.
func (s *BaseTSqlParserListener) ExitClass_type(ctx *Class_typeContext) {}

// EnterClass_type_for_sql_database is called when production class_type_for_sql_database is entered.
func (s *BaseTSqlParserListener) EnterClass_type_for_sql_database(ctx *Class_type_for_sql_databaseContext) {
}

// ExitClass_type_for_sql_database is called when production class_type_for_sql_database is exited.
func (s *BaseTSqlParserListener) ExitClass_type_for_sql_database(ctx *Class_type_for_sql_databaseContext) {
}

// EnterClass_type_for_azure_dw is called when production class_type_for_azure_dw is entered.
func (s *BaseTSqlParserListener) EnterClass_type_for_azure_dw(ctx *Class_type_for_azure_dwContext) {}

// ExitClass_type_for_azure_dw is called when production class_type_for_azure_dw is exited.
func (s *BaseTSqlParserListener) ExitClass_type_for_azure_dw(ctx *Class_type_for_azure_dwContext) {}

// EnterClass_type_for_parallel_dw is called when production class_type_for_parallel_dw is entered.
func (s *BaseTSqlParserListener) EnterClass_type_for_parallel_dw(ctx *Class_type_for_parallel_dwContext) {
}

// ExitClass_type_for_parallel_dw is called when production class_type_for_parallel_dw is exited.
func (s *BaseTSqlParserListener) ExitClass_type_for_parallel_dw(ctx *Class_type_for_parallel_dwContext) {
}

// EnterClass_type_for_grant is called when production class_type_for_grant is entered.
func (s *BaseTSqlParserListener) EnterClass_type_for_grant(ctx *Class_type_for_grantContext) {}

// ExitClass_type_for_grant is called when production class_type_for_grant is exited.
func (s *BaseTSqlParserListener) ExitClass_type_for_grant(ctx *Class_type_for_grantContext) {}

// EnterDrop_availability_group is called when production drop_availability_group is entered.
func (s *BaseTSqlParserListener) EnterDrop_availability_group(ctx *Drop_availability_groupContext) {}

// ExitDrop_availability_group is called when production drop_availability_group is exited.
func (s *BaseTSqlParserListener) ExitDrop_availability_group(ctx *Drop_availability_groupContext) {}

// EnterAlter_availability_group is called when production alter_availability_group is entered.
func (s *BaseTSqlParserListener) EnterAlter_availability_group(ctx *Alter_availability_groupContext) {
}

// ExitAlter_availability_group is called when production alter_availability_group is exited.
func (s *BaseTSqlParserListener) ExitAlter_availability_group(ctx *Alter_availability_groupContext) {}

// EnterAlter_availability_group_start is called when production alter_availability_group_start is entered.
func (s *BaseTSqlParserListener) EnterAlter_availability_group_start(ctx *Alter_availability_group_startContext) {
}

// ExitAlter_availability_group_start is called when production alter_availability_group_start is exited.
func (s *BaseTSqlParserListener) ExitAlter_availability_group_start(ctx *Alter_availability_group_startContext) {
}

// EnterAlter_availability_group_options is called when production alter_availability_group_options is entered.
func (s *BaseTSqlParserListener) EnterAlter_availability_group_options(ctx *Alter_availability_group_optionsContext) {
}

// ExitAlter_availability_group_options is called when production alter_availability_group_options is exited.
func (s *BaseTSqlParserListener) ExitAlter_availability_group_options(ctx *Alter_availability_group_optionsContext) {
}

// EnterIp_v4_failover is called when production ip_v4_failover is entered.
func (s *BaseTSqlParserListener) EnterIp_v4_failover(ctx *Ip_v4_failoverContext) {}

// ExitIp_v4_failover is called when production ip_v4_failover is exited.
func (s *BaseTSqlParserListener) ExitIp_v4_failover(ctx *Ip_v4_failoverContext) {}

// EnterIp_v6_failover is called when production ip_v6_failover is entered.
func (s *BaseTSqlParserListener) EnterIp_v6_failover(ctx *Ip_v6_failoverContext) {}

// ExitIp_v6_failover is called when production ip_v6_failover is exited.
func (s *BaseTSqlParserListener) ExitIp_v6_failover(ctx *Ip_v6_failoverContext) {}

// EnterCreate_or_alter_broker_priority is called when production create_or_alter_broker_priority is entered.
func (s *BaseTSqlParserListener) EnterCreate_or_alter_broker_priority(ctx *Create_or_alter_broker_priorityContext) {
}

// ExitCreate_or_alter_broker_priority is called when production create_or_alter_broker_priority is exited.
func (s *BaseTSqlParserListener) ExitCreate_or_alter_broker_priority(ctx *Create_or_alter_broker_priorityContext) {
}

// EnterDrop_broker_priority is called when production drop_broker_priority is entered.
func (s *BaseTSqlParserListener) EnterDrop_broker_priority(ctx *Drop_broker_priorityContext) {}

// ExitDrop_broker_priority is called when production drop_broker_priority is exited.
func (s *BaseTSqlParserListener) ExitDrop_broker_priority(ctx *Drop_broker_priorityContext) {}

// EnterAlter_certificate is called when production alter_certificate is entered.
func (s *BaseTSqlParserListener) EnterAlter_certificate(ctx *Alter_certificateContext) {}

// ExitAlter_certificate is called when production alter_certificate is exited.
func (s *BaseTSqlParserListener) ExitAlter_certificate(ctx *Alter_certificateContext) {}

// EnterAlter_column_encryption_key is called when production alter_column_encryption_key is entered.
func (s *BaseTSqlParserListener) EnterAlter_column_encryption_key(ctx *Alter_column_encryption_keyContext) {
}

// ExitAlter_column_encryption_key is called when production alter_column_encryption_key is exited.
func (s *BaseTSqlParserListener) ExitAlter_column_encryption_key(ctx *Alter_column_encryption_keyContext) {
}

// EnterCreate_column_encryption_key is called when production create_column_encryption_key is entered.
func (s *BaseTSqlParserListener) EnterCreate_column_encryption_key(ctx *Create_column_encryption_keyContext) {
}

// ExitCreate_column_encryption_key is called when production create_column_encryption_key is exited.
func (s *BaseTSqlParserListener) ExitCreate_column_encryption_key(ctx *Create_column_encryption_keyContext) {
}

// EnterDrop_certificate is called when production drop_certificate is entered.
func (s *BaseTSqlParserListener) EnterDrop_certificate(ctx *Drop_certificateContext) {}

// ExitDrop_certificate is called when production drop_certificate is exited.
func (s *BaseTSqlParserListener) ExitDrop_certificate(ctx *Drop_certificateContext) {}

// EnterDrop_column_encryption_key is called when production drop_column_encryption_key is entered.
func (s *BaseTSqlParserListener) EnterDrop_column_encryption_key(ctx *Drop_column_encryption_keyContext) {
}

// ExitDrop_column_encryption_key is called when production drop_column_encryption_key is exited.
func (s *BaseTSqlParserListener) ExitDrop_column_encryption_key(ctx *Drop_column_encryption_keyContext) {
}

// EnterDrop_column_master_key is called when production drop_column_master_key is entered.
func (s *BaseTSqlParserListener) EnterDrop_column_master_key(ctx *Drop_column_master_keyContext) {}

// ExitDrop_column_master_key is called when production drop_column_master_key is exited.
func (s *BaseTSqlParserListener) ExitDrop_column_master_key(ctx *Drop_column_master_keyContext) {}

// EnterDrop_contract is called when production drop_contract is entered.
func (s *BaseTSqlParserListener) EnterDrop_contract(ctx *Drop_contractContext) {}

// ExitDrop_contract is called when production drop_contract is exited.
func (s *BaseTSqlParserListener) ExitDrop_contract(ctx *Drop_contractContext) {}

// EnterDrop_credential is called when production drop_credential is entered.
func (s *BaseTSqlParserListener) EnterDrop_credential(ctx *Drop_credentialContext) {}

// ExitDrop_credential is called when production drop_credential is exited.
func (s *BaseTSqlParserListener) ExitDrop_credential(ctx *Drop_credentialContext) {}

// EnterDrop_cryptograhic_provider is called when production drop_cryptograhic_provider is entered.
func (s *BaseTSqlParserListener) EnterDrop_cryptograhic_provider(ctx *Drop_cryptograhic_providerContext) {
}

// ExitDrop_cryptograhic_provider is called when production drop_cryptograhic_provider is exited.
func (s *BaseTSqlParserListener) ExitDrop_cryptograhic_provider(ctx *Drop_cryptograhic_providerContext) {
}

// EnterDrop_database is called when production drop_database is entered.
func (s *BaseTSqlParserListener) EnterDrop_database(ctx *Drop_databaseContext) {}

// ExitDrop_database is called when production drop_database is exited.
func (s *BaseTSqlParserListener) ExitDrop_database(ctx *Drop_databaseContext) {}

// EnterDrop_database_audit_specification is called when production drop_database_audit_specification is entered.
func (s *BaseTSqlParserListener) EnterDrop_database_audit_specification(ctx *Drop_database_audit_specificationContext) {
}

// ExitDrop_database_audit_specification is called when production drop_database_audit_specification is exited.
func (s *BaseTSqlParserListener) ExitDrop_database_audit_specification(ctx *Drop_database_audit_specificationContext) {
}

// EnterDrop_database_encryption_key is called when production drop_database_encryption_key is entered.
func (s *BaseTSqlParserListener) EnterDrop_database_encryption_key(ctx *Drop_database_encryption_keyContext) {
}

// ExitDrop_database_encryption_key is called when production drop_database_encryption_key is exited.
func (s *BaseTSqlParserListener) ExitDrop_database_encryption_key(ctx *Drop_database_encryption_keyContext) {
}

// EnterDrop_database_scoped_credential is called when production drop_database_scoped_credential is entered.
func (s *BaseTSqlParserListener) EnterDrop_database_scoped_credential(ctx *Drop_database_scoped_credentialContext) {
}

// ExitDrop_database_scoped_credential is called when production drop_database_scoped_credential is exited.
func (s *BaseTSqlParserListener) ExitDrop_database_scoped_credential(ctx *Drop_database_scoped_credentialContext) {
}

// EnterDrop_default is called when production drop_default is entered.
func (s *BaseTSqlParserListener) EnterDrop_default(ctx *Drop_defaultContext) {}

// ExitDrop_default is called when production drop_default is exited.
func (s *BaseTSqlParserListener) ExitDrop_default(ctx *Drop_defaultContext) {}

// EnterDrop_endpoint is called when production drop_endpoint is entered.
func (s *BaseTSqlParserListener) EnterDrop_endpoint(ctx *Drop_endpointContext) {}

// ExitDrop_endpoint is called when production drop_endpoint is exited.
func (s *BaseTSqlParserListener) ExitDrop_endpoint(ctx *Drop_endpointContext) {}

// EnterDrop_external_data_source is called when production drop_external_data_source is entered.
func (s *BaseTSqlParserListener) EnterDrop_external_data_source(ctx *Drop_external_data_sourceContext) {
}

// ExitDrop_external_data_source is called when production drop_external_data_source is exited.
func (s *BaseTSqlParserListener) ExitDrop_external_data_source(ctx *Drop_external_data_sourceContext) {
}

// EnterDrop_external_file_format is called when production drop_external_file_format is entered.
func (s *BaseTSqlParserListener) EnterDrop_external_file_format(ctx *Drop_external_file_formatContext) {
}

// ExitDrop_external_file_format is called when production drop_external_file_format is exited.
func (s *BaseTSqlParserListener) ExitDrop_external_file_format(ctx *Drop_external_file_formatContext) {
}

// EnterDrop_external_library is called when production drop_external_library is entered.
func (s *BaseTSqlParserListener) EnterDrop_external_library(ctx *Drop_external_libraryContext) {}

// ExitDrop_external_library is called when production drop_external_library is exited.
func (s *BaseTSqlParserListener) ExitDrop_external_library(ctx *Drop_external_libraryContext) {}

// EnterDrop_external_resource_pool is called when production drop_external_resource_pool is entered.
func (s *BaseTSqlParserListener) EnterDrop_external_resource_pool(ctx *Drop_external_resource_poolContext) {
}

// ExitDrop_external_resource_pool is called when production drop_external_resource_pool is exited.
func (s *BaseTSqlParserListener) ExitDrop_external_resource_pool(ctx *Drop_external_resource_poolContext) {
}

// EnterDrop_external_table is called when production drop_external_table is entered.
func (s *BaseTSqlParserListener) EnterDrop_external_table(ctx *Drop_external_tableContext) {}

// ExitDrop_external_table is called when production drop_external_table is exited.
func (s *BaseTSqlParserListener) ExitDrop_external_table(ctx *Drop_external_tableContext) {}

// EnterDrop_event_notifications is called when production drop_event_notifications is entered.
func (s *BaseTSqlParserListener) EnterDrop_event_notifications(ctx *Drop_event_notificationsContext) {
}

// ExitDrop_event_notifications is called when production drop_event_notifications is exited.
func (s *BaseTSqlParserListener) ExitDrop_event_notifications(ctx *Drop_event_notificationsContext) {}

// EnterDrop_event_session is called when production drop_event_session is entered.
func (s *BaseTSqlParserListener) EnterDrop_event_session(ctx *Drop_event_sessionContext) {}

// ExitDrop_event_session is called when production drop_event_session is exited.
func (s *BaseTSqlParserListener) ExitDrop_event_session(ctx *Drop_event_sessionContext) {}

// EnterDrop_fulltext_catalog is called when production drop_fulltext_catalog is entered.
func (s *BaseTSqlParserListener) EnterDrop_fulltext_catalog(ctx *Drop_fulltext_catalogContext) {}

// ExitDrop_fulltext_catalog is called when production drop_fulltext_catalog is exited.
func (s *BaseTSqlParserListener) ExitDrop_fulltext_catalog(ctx *Drop_fulltext_catalogContext) {}

// EnterDrop_fulltext_index is called when production drop_fulltext_index is entered.
func (s *BaseTSqlParserListener) EnterDrop_fulltext_index(ctx *Drop_fulltext_indexContext) {}

// ExitDrop_fulltext_index is called when production drop_fulltext_index is exited.
func (s *BaseTSqlParserListener) ExitDrop_fulltext_index(ctx *Drop_fulltext_indexContext) {}

// EnterDrop_fulltext_stoplist is called when production drop_fulltext_stoplist is entered.
func (s *BaseTSqlParserListener) EnterDrop_fulltext_stoplist(ctx *Drop_fulltext_stoplistContext) {}

// ExitDrop_fulltext_stoplist is called when production drop_fulltext_stoplist is exited.
func (s *BaseTSqlParserListener) ExitDrop_fulltext_stoplist(ctx *Drop_fulltext_stoplistContext) {}

// EnterDrop_login is called when production drop_login is entered.
func (s *BaseTSqlParserListener) EnterDrop_login(ctx *Drop_loginContext) {}

// ExitDrop_login is called when production drop_login is exited.
func (s *BaseTSqlParserListener) ExitDrop_login(ctx *Drop_loginContext) {}

// EnterDrop_master_key is called when production drop_master_key is entered.
func (s *BaseTSqlParserListener) EnterDrop_master_key(ctx *Drop_master_keyContext) {}

// ExitDrop_master_key is called when production drop_master_key is exited.
func (s *BaseTSqlParserListener) ExitDrop_master_key(ctx *Drop_master_keyContext) {}

// EnterDrop_message_type is called when production drop_message_type is entered.
func (s *BaseTSqlParserListener) EnterDrop_message_type(ctx *Drop_message_typeContext) {}

// ExitDrop_message_type is called when production drop_message_type is exited.
func (s *BaseTSqlParserListener) ExitDrop_message_type(ctx *Drop_message_typeContext) {}

// EnterDrop_partition_function is called when production drop_partition_function is entered.
func (s *BaseTSqlParserListener) EnterDrop_partition_function(ctx *Drop_partition_functionContext) {}

// ExitDrop_partition_function is called when production drop_partition_function is exited.
func (s *BaseTSqlParserListener) ExitDrop_partition_function(ctx *Drop_partition_functionContext) {}

// EnterDrop_partition_scheme is called when production drop_partition_scheme is entered.
func (s *BaseTSqlParserListener) EnterDrop_partition_scheme(ctx *Drop_partition_schemeContext) {}

// ExitDrop_partition_scheme is called when production drop_partition_scheme is exited.
func (s *BaseTSqlParserListener) ExitDrop_partition_scheme(ctx *Drop_partition_schemeContext) {}

// EnterDrop_queue is called when production drop_queue is entered.
func (s *BaseTSqlParserListener) EnterDrop_queue(ctx *Drop_queueContext) {}

// ExitDrop_queue is called when production drop_queue is exited.
func (s *BaseTSqlParserListener) ExitDrop_queue(ctx *Drop_queueContext) {}

// EnterDrop_remote_service_binding is called when production drop_remote_service_binding is entered.
func (s *BaseTSqlParserListener) EnterDrop_remote_service_binding(ctx *Drop_remote_service_bindingContext) {
}

// ExitDrop_remote_service_binding is called when production drop_remote_service_binding is exited.
func (s *BaseTSqlParserListener) ExitDrop_remote_service_binding(ctx *Drop_remote_service_bindingContext) {
}

// EnterDrop_resource_pool is called when production drop_resource_pool is entered.
func (s *BaseTSqlParserListener) EnterDrop_resource_pool(ctx *Drop_resource_poolContext) {}

// ExitDrop_resource_pool is called when production drop_resource_pool is exited.
func (s *BaseTSqlParserListener) ExitDrop_resource_pool(ctx *Drop_resource_poolContext) {}

// EnterDrop_db_role is called when production drop_db_role is entered.
func (s *BaseTSqlParserListener) EnterDrop_db_role(ctx *Drop_db_roleContext) {}

// ExitDrop_db_role is called when production drop_db_role is exited.
func (s *BaseTSqlParserListener) ExitDrop_db_role(ctx *Drop_db_roleContext) {}

// EnterDrop_route is called when production drop_route is entered.
func (s *BaseTSqlParserListener) EnterDrop_route(ctx *Drop_routeContext) {}

// ExitDrop_route is called when production drop_route is exited.
func (s *BaseTSqlParserListener) ExitDrop_route(ctx *Drop_routeContext) {}

// EnterDrop_rule is called when production drop_rule is entered.
func (s *BaseTSqlParserListener) EnterDrop_rule(ctx *Drop_ruleContext) {}

// ExitDrop_rule is called when production drop_rule is exited.
func (s *BaseTSqlParserListener) ExitDrop_rule(ctx *Drop_ruleContext) {}

// EnterDrop_schema is called when production drop_schema is entered.
func (s *BaseTSqlParserListener) EnterDrop_schema(ctx *Drop_schemaContext) {}

// ExitDrop_schema is called when production drop_schema is exited.
func (s *BaseTSqlParserListener) ExitDrop_schema(ctx *Drop_schemaContext) {}

// EnterDrop_search_property_list is called when production drop_search_property_list is entered.
func (s *BaseTSqlParserListener) EnterDrop_search_property_list(ctx *Drop_search_property_listContext) {
}

// ExitDrop_search_property_list is called when production drop_search_property_list is exited.
func (s *BaseTSqlParserListener) ExitDrop_search_property_list(ctx *Drop_search_property_listContext) {
}

// EnterDrop_security_policy is called when production drop_security_policy is entered.
func (s *BaseTSqlParserListener) EnterDrop_security_policy(ctx *Drop_security_policyContext) {}

// ExitDrop_security_policy is called when production drop_security_policy is exited.
func (s *BaseTSqlParserListener) ExitDrop_security_policy(ctx *Drop_security_policyContext) {}

// EnterDrop_sequence is called when production drop_sequence is entered.
func (s *BaseTSqlParserListener) EnterDrop_sequence(ctx *Drop_sequenceContext) {}

// ExitDrop_sequence is called when production drop_sequence is exited.
func (s *BaseTSqlParserListener) ExitDrop_sequence(ctx *Drop_sequenceContext) {}

// EnterDrop_server_audit is called when production drop_server_audit is entered.
func (s *BaseTSqlParserListener) EnterDrop_server_audit(ctx *Drop_server_auditContext) {}

// ExitDrop_server_audit is called when production drop_server_audit is exited.
func (s *BaseTSqlParserListener) ExitDrop_server_audit(ctx *Drop_server_auditContext) {}

// EnterDrop_server_audit_specification is called when production drop_server_audit_specification is entered.
func (s *BaseTSqlParserListener) EnterDrop_server_audit_specification(ctx *Drop_server_audit_specificationContext) {
}

// ExitDrop_server_audit_specification is called when production drop_server_audit_specification is exited.
func (s *BaseTSqlParserListener) ExitDrop_server_audit_specification(ctx *Drop_server_audit_specificationContext) {
}

// EnterDrop_server_role is called when production drop_server_role is entered.
func (s *BaseTSqlParserListener) EnterDrop_server_role(ctx *Drop_server_roleContext) {}

// ExitDrop_server_role is called when production drop_server_role is exited.
func (s *BaseTSqlParserListener) ExitDrop_server_role(ctx *Drop_server_roleContext) {}

// EnterDrop_service is called when production drop_service is entered.
func (s *BaseTSqlParserListener) EnterDrop_service(ctx *Drop_serviceContext) {}

// ExitDrop_service is called when production drop_service is exited.
func (s *BaseTSqlParserListener) ExitDrop_service(ctx *Drop_serviceContext) {}

// EnterDrop_signature is called when production drop_signature is entered.
func (s *BaseTSqlParserListener) EnterDrop_signature(ctx *Drop_signatureContext) {}

// ExitDrop_signature is called when production drop_signature is exited.
func (s *BaseTSqlParserListener) ExitDrop_signature(ctx *Drop_signatureContext) {}

// EnterDrop_statistics_name_azure_dw_and_pdw is called when production drop_statistics_name_azure_dw_and_pdw is entered.
func (s *BaseTSqlParserListener) EnterDrop_statistics_name_azure_dw_and_pdw(ctx *Drop_statistics_name_azure_dw_and_pdwContext) {
}

// ExitDrop_statistics_name_azure_dw_and_pdw is called when production drop_statistics_name_azure_dw_and_pdw is exited.
func (s *BaseTSqlParserListener) ExitDrop_statistics_name_azure_dw_and_pdw(ctx *Drop_statistics_name_azure_dw_and_pdwContext) {
}

// EnterDrop_symmetric_key is called when production drop_symmetric_key is entered.
func (s *BaseTSqlParserListener) EnterDrop_symmetric_key(ctx *Drop_symmetric_keyContext) {}

// ExitDrop_symmetric_key is called when production drop_symmetric_key is exited.
func (s *BaseTSqlParserListener) ExitDrop_symmetric_key(ctx *Drop_symmetric_keyContext) {}

// EnterDrop_synonym is called when production drop_synonym is entered.
func (s *BaseTSqlParserListener) EnterDrop_synonym(ctx *Drop_synonymContext) {}

// ExitDrop_synonym is called when production drop_synonym is exited.
func (s *BaseTSqlParserListener) ExitDrop_synonym(ctx *Drop_synonymContext) {}

// EnterDrop_user is called when production drop_user is entered.
func (s *BaseTSqlParserListener) EnterDrop_user(ctx *Drop_userContext) {}

// ExitDrop_user is called when production drop_user is exited.
func (s *BaseTSqlParserListener) ExitDrop_user(ctx *Drop_userContext) {}

// EnterDrop_workload_group is called when production drop_workload_group is entered.
func (s *BaseTSqlParserListener) EnterDrop_workload_group(ctx *Drop_workload_groupContext) {}

// ExitDrop_workload_group is called when production drop_workload_group is exited.
func (s *BaseTSqlParserListener) ExitDrop_workload_group(ctx *Drop_workload_groupContext) {}

// EnterDrop_xml_schema_collection is called when production drop_xml_schema_collection is entered.
func (s *BaseTSqlParserListener) EnterDrop_xml_schema_collection(ctx *Drop_xml_schema_collectionContext) {
}

// ExitDrop_xml_schema_collection is called when production drop_xml_schema_collection is exited.
func (s *BaseTSqlParserListener) ExitDrop_xml_schema_collection(ctx *Drop_xml_schema_collectionContext) {
}

// EnterDisable_trigger is called when production disable_trigger is entered.
func (s *BaseTSqlParserListener) EnterDisable_trigger(ctx *Disable_triggerContext) {}

// ExitDisable_trigger is called when production disable_trigger is exited.
func (s *BaseTSqlParserListener) ExitDisable_trigger(ctx *Disable_triggerContext) {}

// EnterEnable_trigger is called when production enable_trigger is entered.
func (s *BaseTSqlParserListener) EnterEnable_trigger(ctx *Enable_triggerContext) {}

// ExitEnable_trigger is called when production enable_trigger is exited.
func (s *BaseTSqlParserListener) ExitEnable_trigger(ctx *Enable_triggerContext) {}

// EnterLock_table is called when production lock_table is entered.
func (s *BaseTSqlParserListener) EnterLock_table(ctx *Lock_tableContext) {}

// ExitLock_table is called when production lock_table is exited.
func (s *BaseTSqlParserListener) ExitLock_table(ctx *Lock_tableContext) {}

// EnterTruncate_table is called when production truncate_table is entered.
func (s *BaseTSqlParserListener) EnterTruncate_table(ctx *Truncate_tableContext) {}

// ExitTruncate_table is called when production truncate_table is exited.
func (s *BaseTSqlParserListener) ExitTruncate_table(ctx *Truncate_tableContext) {}

// EnterCreate_column_master_key is called when production create_column_master_key is entered.
func (s *BaseTSqlParserListener) EnterCreate_column_master_key(ctx *Create_column_master_keyContext) {
}

// ExitCreate_column_master_key is called when production create_column_master_key is exited.
func (s *BaseTSqlParserListener) ExitCreate_column_master_key(ctx *Create_column_master_keyContext) {}

// EnterAlter_credential is called when production alter_credential is entered.
func (s *BaseTSqlParserListener) EnterAlter_credential(ctx *Alter_credentialContext) {}

// ExitAlter_credential is called when production alter_credential is exited.
func (s *BaseTSqlParserListener) ExitAlter_credential(ctx *Alter_credentialContext) {}

// EnterCreate_credential is called when production create_credential is entered.
func (s *BaseTSqlParserListener) EnterCreate_credential(ctx *Create_credentialContext) {}

// ExitCreate_credential is called when production create_credential is exited.
func (s *BaseTSqlParserListener) ExitCreate_credential(ctx *Create_credentialContext) {}

// EnterAlter_cryptographic_provider is called when production alter_cryptographic_provider is entered.
func (s *BaseTSqlParserListener) EnterAlter_cryptographic_provider(ctx *Alter_cryptographic_providerContext) {
}

// ExitAlter_cryptographic_provider is called when production alter_cryptographic_provider is exited.
func (s *BaseTSqlParserListener) ExitAlter_cryptographic_provider(ctx *Alter_cryptographic_providerContext) {
}

// EnterCreate_cryptographic_provider is called when production create_cryptographic_provider is entered.
func (s *BaseTSqlParserListener) EnterCreate_cryptographic_provider(ctx *Create_cryptographic_providerContext) {
}

// ExitCreate_cryptographic_provider is called when production create_cryptographic_provider is exited.
func (s *BaseTSqlParserListener) ExitCreate_cryptographic_provider(ctx *Create_cryptographic_providerContext) {
}

// EnterCreate_event_notification is called when production create_event_notification is entered.
func (s *BaseTSqlParserListener) EnterCreate_event_notification(ctx *Create_event_notificationContext) {
}

// ExitCreate_event_notification is called when production create_event_notification is exited.
func (s *BaseTSqlParserListener) ExitCreate_event_notification(ctx *Create_event_notificationContext) {
}

// EnterCreate_or_alter_event_session is called when production create_or_alter_event_session is entered.
func (s *BaseTSqlParserListener) EnterCreate_or_alter_event_session(ctx *Create_or_alter_event_sessionContext) {
}

// ExitCreate_or_alter_event_session is called when production create_or_alter_event_session is exited.
func (s *BaseTSqlParserListener) ExitCreate_or_alter_event_session(ctx *Create_or_alter_event_sessionContext) {
}

// EnterEvent_session_predicate_expression is called when production event_session_predicate_expression is entered.
func (s *BaseTSqlParserListener) EnterEvent_session_predicate_expression(ctx *Event_session_predicate_expressionContext) {
}

// ExitEvent_session_predicate_expression is called when production event_session_predicate_expression is exited.
func (s *BaseTSqlParserListener) ExitEvent_session_predicate_expression(ctx *Event_session_predicate_expressionContext) {
}

// EnterEvent_session_predicate_factor is called when production event_session_predicate_factor is entered.
func (s *BaseTSqlParserListener) EnterEvent_session_predicate_factor(ctx *Event_session_predicate_factorContext) {
}

// ExitEvent_session_predicate_factor is called when production event_session_predicate_factor is exited.
func (s *BaseTSqlParserListener) ExitEvent_session_predicate_factor(ctx *Event_session_predicate_factorContext) {
}

// EnterEvent_session_predicate_leaf is called when production event_session_predicate_leaf is entered.
func (s *BaseTSqlParserListener) EnterEvent_session_predicate_leaf(ctx *Event_session_predicate_leafContext) {
}

// ExitEvent_session_predicate_leaf is called when production event_session_predicate_leaf is exited.
func (s *BaseTSqlParserListener) ExitEvent_session_predicate_leaf(ctx *Event_session_predicate_leafContext) {
}

// EnterAlter_external_data_source is called when production alter_external_data_source is entered.
func (s *BaseTSqlParserListener) EnterAlter_external_data_source(ctx *Alter_external_data_sourceContext) {
}

// ExitAlter_external_data_source is called when production alter_external_data_source is exited.
func (s *BaseTSqlParserListener) ExitAlter_external_data_source(ctx *Alter_external_data_sourceContext) {
}

// EnterAlter_external_library is called when production alter_external_library is entered.
func (s *BaseTSqlParserListener) EnterAlter_external_library(ctx *Alter_external_libraryContext) {}

// ExitAlter_external_library is called when production alter_external_library is exited.
func (s *BaseTSqlParserListener) ExitAlter_external_library(ctx *Alter_external_libraryContext) {}

// EnterCreate_external_library is called when production create_external_library is entered.
func (s *BaseTSqlParserListener) EnterCreate_external_library(ctx *Create_external_libraryContext) {}

// ExitCreate_external_library is called when production create_external_library is exited.
func (s *BaseTSqlParserListener) ExitCreate_external_library(ctx *Create_external_libraryContext) {}

// EnterAlter_external_resource_pool is called when production alter_external_resource_pool is entered.
func (s *BaseTSqlParserListener) EnterAlter_external_resource_pool(ctx *Alter_external_resource_poolContext) {
}

// ExitAlter_external_resource_pool is called when production alter_external_resource_pool is exited.
func (s *BaseTSqlParserListener) ExitAlter_external_resource_pool(ctx *Alter_external_resource_poolContext) {
}

// EnterCreate_external_resource_pool is called when production create_external_resource_pool is entered.
func (s *BaseTSqlParserListener) EnterCreate_external_resource_pool(ctx *Create_external_resource_poolContext) {
}

// ExitCreate_external_resource_pool is called when production create_external_resource_pool is exited.
func (s *BaseTSqlParserListener) ExitCreate_external_resource_pool(ctx *Create_external_resource_poolContext) {
}

// EnterAlter_fulltext_catalog is called when production alter_fulltext_catalog is entered.
func (s *BaseTSqlParserListener) EnterAlter_fulltext_catalog(ctx *Alter_fulltext_catalogContext) {}

// ExitAlter_fulltext_catalog is called when production alter_fulltext_catalog is exited.
func (s *BaseTSqlParserListener) ExitAlter_fulltext_catalog(ctx *Alter_fulltext_catalogContext) {}

// EnterCreate_fulltext_catalog is called when production create_fulltext_catalog is entered.
func (s *BaseTSqlParserListener) EnterCreate_fulltext_catalog(ctx *Create_fulltext_catalogContext) {}

// ExitCreate_fulltext_catalog is called when production create_fulltext_catalog is exited.
func (s *BaseTSqlParserListener) ExitCreate_fulltext_catalog(ctx *Create_fulltext_catalogContext) {}

// EnterAlter_fulltext_stoplist is called when production alter_fulltext_stoplist is entered.
func (s *BaseTSqlParserListener) EnterAlter_fulltext_stoplist(ctx *Alter_fulltext_stoplistContext) {}

// ExitAlter_fulltext_stoplist is called when production alter_fulltext_stoplist is exited.
func (s *BaseTSqlParserListener) ExitAlter_fulltext_stoplist(ctx *Alter_fulltext_stoplistContext) {}

// EnterCreate_fulltext_stoplist is called when production create_fulltext_stoplist is entered.
func (s *BaseTSqlParserListener) EnterCreate_fulltext_stoplist(ctx *Create_fulltext_stoplistContext) {
}

// ExitCreate_fulltext_stoplist is called when production create_fulltext_stoplist is exited.
func (s *BaseTSqlParserListener) ExitCreate_fulltext_stoplist(ctx *Create_fulltext_stoplistContext) {}

// EnterAlter_login_sql_server is called when production alter_login_sql_server is entered.
func (s *BaseTSqlParserListener) EnterAlter_login_sql_server(ctx *Alter_login_sql_serverContext) {}

// ExitAlter_login_sql_server is called when production alter_login_sql_server is exited.
func (s *BaseTSqlParserListener) ExitAlter_login_sql_server(ctx *Alter_login_sql_serverContext) {}

// EnterCreate_login_sql_server is called when production create_login_sql_server is entered.
func (s *BaseTSqlParserListener) EnterCreate_login_sql_server(ctx *Create_login_sql_serverContext) {}

// ExitCreate_login_sql_server is called when production create_login_sql_server is exited.
func (s *BaseTSqlParserListener) ExitCreate_login_sql_server(ctx *Create_login_sql_serverContext) {}

// EnterAlter_login_azure_sql is called when production alter_login_azure_sql is entered.
func (s *BaseTSqlParserListener) EnterAlter_login_azure_sql(ctx *Alter_login_azure_sqlContext) {}

// ExitAlter_login_azure_sql is called when production alter_login_azure_sql is exited.
func (s *BaseTSqlParserListener) ExitAlter_login_azure_sql(ctx *Alter_login_azure_sqlContext) {}

// EnterCreate_login_azure_sql is called when production create_login_azure_sql is entered.
func (s *BaseTSqlParserListener) EnterCreate_login_azure_sql(ctx *Create_login_azure_sqlContext) {}

// ExitCreate_login_azure_sql is called when production create_login_azure_sql is exited.
func (s *BaseTSqlParserListener) ExitCreate_login_azure_sql(ctx *Create_login_azure_sqlContext) {}

// EnterAlter_login_azure_sql_dw_and_pdw is called when production alter_login_azure_sql_dw_and_pdw is entered.
func (s *BaseTSqlParserListener) EnterAlter_login_azure_sql_dw_and_pdw(ctx *Alter_login_azure_sql_dw_and_pdwContext) {
}

// ExitAlter_login_azure_sql_dw_and_pdw is called when production alter_login_azure_sql_dw_and_pdw is exited.
func (s *BaseTSqlParserListener) ExitAlter_login_azure_sql_dw_and_pdw(ctx *Alter_login_azure_sql_dw_and_pdwContext) {
}

// EnterCreate_login_pdw is called when production create_login_pdw is entered.
func (s *BaseTSqlParserListener) EnterCreate_login_pdw(ctx *Create_login_pdwContext) {}

// ExitCreate_login_pdw is called when production create_login_pdw is exited.
func (s *BaseTSqlParserListener) ExitCreate_login_pdw(ctx *Create_login_pdwContext) {}

// EnterAlter_master_key_sql_server is called when production alter_master_key_sql_server is entered.
func (s *BaseTSqlParserListener) EnterAlter_master_key_sql_server(ctx *Alter_master_key_sql_serverContext) {
}

// ExitAlter_master_key_sql_server is called when production alter_master_key_sql_server is exited.
func (s *BaseTSqlParserListener) ExitAlter_master_key_sql_server(ctx *Alter_master_key_sql_serverContext) {
}

// EnterCreate_master_key_sql_server is called when production create_master_key_sql_server is entered.
func (s *BaseTSqlParserListener) EnterCreate_master_key_sql_server(ctx *Create_master_key_sql_serverContext) {
}

// ExitCreate_master_key_sql_server is called when production create_master_key_sql_server is exited.
func (s *BaseTSqlParserListener) ExitCreate_master_key_sql_server(ctx *Create_master_key_sql_serverContext) {
}

// EnterAlter_master_key_azure_sql is called when production alter_master_key_azure_sql is entered.
func (s *BaseTSqlParserListener) EnterAlter_master_key_azure_sql(ctx *Alter_master_key_azure_sqlContext) {
}

// ExitAlter_master_key_azure_sql is called when production alter_master_key_azure_sql is exited.
func (s *BaseTSqlParserListener) ExitAlter_master_key_azure_sql(ctx *Alter_master_key_azure_sqlContext) {
}

// EnterCreate_master_key_azure_sql is called when production create_master_key_azure_sql is entered.
func (s *BaseTSqlParserListener) EnterCreate_master_key_azure_sql(ctx *Create_master_key_azure_sqlContext) {
}

// ExitCreate_master_key_azure_sql is called when production create_master_key_azure_sql is exited.
func (s *BaseTSqlParserListener) ExitCreate_master_key_azure_sql(ctx *Create_master_key_azure_sqlContext) {
}

// EnterAlter_message_type is called when production alter_message_type is entered.
func (s *BaseTSqlParserListener) EnterAlter_message_type(ctx *Alter_message_typeContext) {}

// ExitAlter_message_type is called when production alter_message_type is exited.
func (s *BaseTSqlParserListener) ExitAlter_message_type(ctx *Alter_message_typeContext) {}

// EnterAlter_partition_function is called when production alter_partition_function is entered.
func (s *BaseTSqlParserListener) EnterAlter_partition_function(ctx *Alter_partition_functionContext) {
}

// ExitAlter_partition_function is called when production alter_partition_function is exited.
func (s *BaseTSqlParserListener) ExitAlter_partition_function(ctx *Alter_partition_functionContext) {}

// EnterAlter_partition_scheme is called when production alter_partition_scheme is entered.
func (s *BaseTSqlParserListener) EnterAlter_partition_scheme(ctx *Alter_partition_schemeContext) {}

// ExitAlter_partition_scheme is called when production alter_partition_scheme is exited.
func (s *BaseTSqlParserListener) ExitAlter_partition_scheme(ctx *Alter_partition_schemeContext) {}

// EnterAlter_remote_service_binding is called when production alter_remote_service_binding is entered.
func (s *BaseTSqlParserListener) EnterAlter_remote_service_binding(ctx *Alter_remote_service_bindingContext) {
}

// ExitAlter_remote_service_binding is called when production alter_remote_service_binding is exited.
func (s *BaseTSqlParserListener) ExitAlter_remote_service_binding(ctx *Alter_remote_service_bindingContext) {
}

// EnterCreate_remote_service_binding is called when production create_remote_service_binding is entered.
func (s *BaseTSqlParserListener) EnterCreate_remote_service_binding(ctx *Create_remote_service_bindingContext) {
}

// ExitCreate_remote_service_binding is called when production create_remote_service_binding is exited.
func (s *BaseTSqlParserListener) ExitCreate_remote_service_binding(ctx *Create_remote_service_bindingContext) {
}

// EnterCreate_resource_pool is called when production create_resource_pool is entered.
func (s *BaseTSqlParserListener) EnterCreate_resource_pool(ctx *Create_resource_poolContext) {}

// ExitCreate_resource_pool is called when production create_resource_pool is exited.
func (s *BaseTSqlParserListener) ExitCreate_resource_pool(ctx *Create_resource_poolContext) {}

// EnterAlter_resource_governor is called when production alter_resource_governor is entered.
func (s *BaseTSqlParserListener) EnterAlter_resource_governor(ctx *Alter_resource_governorContext) {}

// ExitAlter_resource_governor is called when production alter_resource_governor is exited.
func (s *BaseTSqlParserListener) ExitAlter_resource_governor(ctx *Alter_resource_governorContext) {}

// EnterAlter_db_role is called when production alter_db_role is entered.
func (s *BaseTSqlParserListener) EnterAlter_db_role(ctx *Alter_db_roleContext) {}

// ExitAlter_db_role is called when production alter_db_role is exited.
func (s *BaseTSqlParserListener) ExitAlter_db_role(ctx *Alter_db_roleContext) {}

// EnterCreate_db_role is called when production create_db_role is entered.
func (s *BaseTSqlParserListener) EnterCreate_db_role(ctx *Create_db_roleContext) {}

// ExitCreate_db_role is called when production create_db_role is exited.
func (s *BaseTSqlParserListener) ExitCreate_db_role(ctx *Create_db_roleContext) {}

// EnterCreate_route is called when production create_route is entered.
func (s *BaseTSqlParserListener) EnterCreate_route(ctx *Create_routeContext) {}

// ExitCreate_route is called when production create_route is exited.
func (s *BaseTSqlParserListener) ExitCreate_route(ctx *Create_routeContext) {}

// EnterCreate_rule is called when production create_rule is entered.
func (s *BaseTSqlParserListener) EnterCreate_rule(ctx *Create_ruleContext) {}

// ExitCreate_rule is called when production create_rule is exited.
func (s *BaseTSqlParserListener) ExitCreate_rule(ctx *Create_ruleContext) {}

// EnterAlter_schema_sql is called when production alter_schema_sql is entered.
func (s *BaseTSqlParserListener) EnterAlter_schema_sql(ctx *Alter_schema_sqlContext) {}

// ExitAlter_schema_sql is called when production alter_schema_sql is exited.
func (s *BaseTSqlParserListener) ExitAlter_schema_sql(ctx *Alter_schema_sqlContext) {}

// EnterCreate_schema is called when production create_schema is entered.
func (s *BaseTSqlParserListener) EnterCreate_schema(ctx *Create_schemaContext) {}

// ExitCreate_schema is called when production create_schema is exited.
func (s *BaseTSqlParserListener) ExitCreate_schema(ctx *Create_schemaContext) {}

// EnterCreate_schema_azure_sql_dw_and_pdw is called when production create_schema_azure_sql_dw_and_pdw is entered.
func (s *BaseTSqlParserListener) EnterCreate_schema_azure_sql_dw_and_pdw(ctx *Create_schema_azure_sql_dw_and_pdwContext) {
}

// ExitCreate_schema_azure_sql_dw_and_pdw is called when production create_schema_azure_sql_dw_and_pdw is exited.
func (s *BaseTSqlParserListener) ExitCreate_schema_azure_sql_dw_and_pdw(ctx *Create_schema_azure_sql_dw_and_pdwContext) {
}

// EnterAlter_schema_azure_sql_dw_and_pdw is called when production alter_schema_azure_sql_dw_and_pdw is entered.
func (s *BaseTSqlParserListener) EnterAlter_schema_azure_sql_dw_and_pdw(ctx *Alter_schema_azure_sql_dw_and_pdwContext) {
}

// ExitAlter_schema_azure_sql_dw_and_pdw is called when production alter_schema_azure_sql_dw_and_pdw is exited.
func (s *BaseTSqlParserListener) ExitAlter_schema_azure_sql_dw_and_pdw(ctx *Alter_schema_azure_sql_dw_and_pdwContext) {
}

// EnterCreate_search_property_list is called when production create_search_property_list is entered.
func (s *BaseTSqlParserListener) EnterCreate_search_property_list(ctx *Create_search_property_listContext) {
}

// ExitCreate_search_property_list is called when production create_search_property_list is exited.
func (s *BaseTSqlParserListener) ExitCreate_search_property_list(ctx *Create_search_property_listContext) {
}

// EnterCreate_security_policy is called when production create_security_policy is entered.
func (s *BaseTSqlParserListener) EnterCreate_security_policy(ctx *Create_security_policyContext) {}

// ExitCreate_security_policy is called when production create_security_policy is exited.
func (s *BaseTSqlParserListener) ExitCreate_security_policy(ctx *Create_security_policyContext) {}

// EnterAlter_sequence is called when production alter_sequence is entered.
func (s *BaseTSqlParserListener) EnterAlter_sequence(ctx *Alter_sequenceContext) {}

// ExitAlter_sequence is called when production alter_sequence is exited.
func (s *BaseTSqlParserListener) ExitAlter_sequence(ctx *Alter_sequenceContext) {}

// EnterCreate_sequence is called when production create_sequence is entered.
func (s *BaseTSqlParserListener) EnterCreate_sequence(ctx *Create_sequenceContext) {}

// ExitCreate_sequence is called when production create_sequence is exited.
func (s *BaseTSqlParserListener) ExitCreate_sequence(ctx *Create_sequenceContext) {}

// EnterAlter_server_audit is called when production alter_server_audit is entered.
func (s *BaseTSqlParserListener) EnterAlter_server_audit(ctx *Alter_server_auditContext) {}

// ExitAlter_server_audit is called when production alter_server_audit is exited.
func (s *BaseTSqlParserListener) ExitAlter_server_audit(ctx *Alter_server_auditContext) {}

// EnterCreate_server_audit is called when production create_server_audit is entered.
func (s *BaseTSqlParserListener) EnterCreate_server_audit(ctx *Create_server_auditContext) {}

// ExitCreate_server_audit is called when production create_server_audit is exited.
func (s *BaseTSqlParserListener) ExitCreate_server_audit(ctx *Create_server_auditContext) {}

// EnterAlter_server_audit_specification is called when production alter_server_audit_specification is entered.
func (s *BaseTSqlParserListener) EnterAlter_server_audit_specification(ctx *Alter_server_audit_specificationContext) {
}

// ExitAlter_server_audit_specification is called when production alter_server_audit_specification is exited.
func (s *BaseTSqlParserListener) ExitAlter_server_audit_specification(ctx *Alter_server_audit_specificationContext) {
}

// EnterCreate_server_audit_specification is called when production create_server_audit_specification is entered.
func (s *BaseTSqlParserListener) EnterCreate_server_audit_specification(ctx *Create_server_audit_specificationContext) {
}

// ExitCreate_server_audit_specification is called when production create_server_audit_specification is exited.
func (s *BaseTSqlParserListener) ExitCreate_server_audit_specification(ctx *Create_server_audit_specificationContext) {
}

// EnterAlter_server_configuration is called when production alter_server_configuration is entered.
func (s *BaseTSqlParserListener) EnterAlter_server_configuration(ctx *Alter_server_configurationContext) {
}

// ExitAlter_server_configuration is called when production alter_server_configuration is exited.
func (s *BaseTSqlParserListener) ExitAlter_server_configuration(ctx *Alter_server_configurationContext) {
}

// EnterAlter_server_role is called when production alter_server_role is entered.
func (s *BaseTSqlParserListener) EnterAlter_server_role(ctx *Alter_server_roleContext) {}

// ExitAlter_server_role is called when production alter_server_role is exited.
func (s *BaseTSqlParserListener) ExitAlter_server_role(ctx *Alter_server_roleContext) {}

// EnterCreate_server_role is called when production create_server_role is entered.
func (s *BaseTSqlParserListener) EnterCreate_server_role(ctx *Create_server_roleContext) {}

// ExitCreate_server_role is called when production create_server_role is exited.
func (s *BaseTSqlParserListener) ExitCreate_server_role(ctx *Create_server_roleContext) {}

// EnterAlter_server_role_pdw is called when production alter_server_role_pdw is entered.
func (s *BaseTSqlParserListener) EnterAlter_server_role_pdw(ctx *Alter_server_role_pdwContext) {}

// ExitAlter_server_role_pdw is called when production alter_server_role_pdw is exited.
func (s *BaseTSqlParserListener) ExitAlter_server_role_pdw(ctx *Alter_server_role_pdwContext) {}

// EnterAlter_service is called when production alter_service is entered.
func (s *BaseTSqlParserListener) EnterAlter_service(ctx *Alter_serviceContext) {}

// ExitAlter_service is called when production alter_service is exited.
func (s *BaseTSqlParserListener) ExitAlter_service(ctx *Alter_serviceContext) {}

// EnterCreate_service is called when production create_service is entered.
func (s *BaseTSqlParserListener) EnterCreate_service(ctx *Create_serviceContext) {}

// ExitCreate_service is called when production create_service is exited.
func (s *BaseTSqlParserListener) ExitCreate_service(ctx *Create_serviceContext) {}

// EnterAlter_service_master_key is called when production alter_service_master_key is entered.
func (s *BaseTSqlParserListener) EnterAlter_service_master_key(ctx *Alter_service_master_keyContext) {
}

// ExitAlter_service_master_key is called when production alter_service_master_key is exited.
func (s *BaseTSqlParserListener) ExitAlter_service_master_key(ctx *Alter_service_master_keyContext) {}

// EnterAlter_symmetric_key is called when production alter_symmetric_key is entered.
func (s *BaseTSqlParserListener) EnterAlter_symmetric_key(ctx *Alter_symmetric_keyContext) {}

// ExitAlter_symmetric_key is called when production alter_symmetric_key is exited.
func (s *BaseTSqlParserListener) ExitAlter_symmetric_key(ctx *Alter_symmetric_keyContext) {}

// EnterCreate_synonym is called when production create_synonym is entered.
func (s *BaseTSqlParserListener) EnterCreate_synonym(ctx *Create_synonymContext) {}

// ExitCreate_synonym is called when production create_synonym is exited.
func (s *BaseTSqlParserListener) ExitCreate_synonym(ctx *Create_synonymContext) {}

// EnterAlter_user is called when production alter_user is entered.
func (s *BaseTSqlParserListener) EnterAlter_user(ctx *Alter_userContext) {}

// ExitAlter_user is called when production alter_user is exited.
func (s *BaseTSqlParserListener) ExitAlter_user(ctx *Alter_userContext) {}

// EnterCreate_user is called when production create_user is entered.
func (s *BaseTSqlParserListener) EnterCreate_user(ctx *Create_userContext) {}

// ExitCreate_user is called when production create_user is exited.
func (s *BaseTSqlParserListener) ExitCreate_user(ctx *Create_userContext) {}

// EnterCreate_user_azure_sql_dw is called when production create_user_azure_sql_dw is entered.
func (s *BaseTSqlParserListener) EnterCreate_user_azure_sql_dw(ctx *Create_user_azure_sql_dwContext) {
}

// ExitCreate_user_azure_sql_dw is called when production create_user_azure_sql_dw is exited.
func (s *BaseTSqlParserListener) ExitCreate_user_azure_sql_dw(ctx *Create_user_azure_sql_dwContext) {}

// EnterAlter_user_azure_sql is called when production alter_user_azure_sql is entered.
func (s *BaseTSqlParserListener) EnterAlter_user_azure_sql(ctx *Alter_user_azure_sqlContext) {}

// ExitAlter_user_azure_sql is called when production alter_user_azure_sql is exited.
func (s *BaseTSqlParserListener) ExitAlter_user_azure_sql(ctx *Alter_user_azure_sqlContext) {}

// EnterAlter_workload_group is called when production alter_workload_group is entered.
func (s *BaseTSqlParserListener) EnterAlter_workload_group(ctx *Alter_workload_groupContext) {}

// ExitAlter_workload_group is called when production alter_workload_group is exited.
func (s *BaseTSqlParserListener) ExitAlter_workload_group(ctx *Alter_workload_groupContext) {}

// EnterCreate_workload_group is called when production create_workload_group is entered.
func (s *BaseTSqlParserListener) EnterCreate_workload_group(ctx *Create_workload_groupContext) {}

// ExitCreate_workload_group is called when production create_workload_group is exited.
func (s *BaseTSqlParserListener) ExitCreate_workload_group(ctx *Create_workload_groupContext) {}

// EnterCreate_xml_schema_collection is called when production create_xml_schema_collection is entered.
func (s *BaseTSqlParserListener) EnterCreate_xml_schema_collection(ctx *Create_xml_schema_collectionContext) {
}

// ExitCreate_xml_schema_collection is called when production create_xml_schema_collection is exited.
func (s *BaseTSqlParserListener) ExitCreate_xml_schema_collection(ctx *Create_xml_schema_collectionContext) {
}

// EnterCreate_partition_function is called when production create_partition_function is entered.
func (s *BaseTSqlParserListener) EnterCreate_partition_function(ctx *Create_partition_functionContext) {
}

// ExitCreate_partition_function is called when production create_partition_function is exited.
func (s *BaseTSqlParserListener) ExitCreate_partition_function(ctx *Create_partition_functionContext) {
}

// EnterCreate_partition_scheme is called when production create_partition_scheme is entered.
func (s *BaseTSqlParserListener) EnterCreate_partition_scheme(ctx *Create_partition_schemeContext) {}

// ExitCreate_partition_scheme is called when production create_partition_scheme is exited.
func (s *BaseTSqlParserListener) ExitCreate_partition_scheme(ctx *Create_partition_schemeContext) {}

// EnterCreate_queue is called when production create_queue is entered.
func (s *BaseTSqlParserListener) EnterCreate_queue(ctx *Create_queueContext) {}

// ExitCreate_queue is called when production create_queue is exited.
func (s *BaseTSqlParserListener) ExitCreate_queue(ctx *Create_queueContext) {}

// EnterQueue_settings is called when production queue_settings is entered.
func (s *BaseTSqlParserListener) EnterQueue_settings(ctx *Queue_settingsContext) {}

// ExitQueue_settings is called when production queue_settings is exited.
func (s *BaseTSqlParserListener) ExitQueue_settings(ctx *Queue_settingsContext) {}

// EnterAlter_queue is called when production alter_queue is entered.
func (s *BaseTSqlParserListener) EnterAlter_queue(ctx *Alter_queueContext) {}

// ExitAlter_queue is called when production alter_queue is exited.
func (s *BaseTSqlParserListener) ExitAlter_queue(ctx *Alter_queueContext) {}

// EnterQueue_action is called when production queue_action is entered.
func (s *BaseTSqlParserListener) EnterQueue_action(ctx *Queue_actionContext) {}

// ExitQueue_action is called when production queue_action is exited.
func (s *BaseTSqlParserListener) ExitQueue_action(ctx *Queue_actionContext) {}

// EnterQueue_rebuild_options is called when production queue_rebuild_options is entered.
func (s *BaseTSqlParserListener) EnterQueue_rebuild_options(ctx *Queue_rebuild_optionsContext) {}

// ExitQueue_rebuild_options is called when production queue_rebuild_options is exited.
func (s *BaseTSqlParserListener) ExitQueue_rebuild_options(ctx *Queue_rebuild_optionsContext) {}

// EnterCreate_contract is called when production create_contract is entered.
func (s *BaseTSqlParserListener) EnterCreate_contract(ctx *Create_contractContext) {}

// ExitCreate_contract is called when production create_contract is exited.
func (s *BaseTSqlParserListener) ExitCreate_contract(ctx *Create_contractContext) {}

// EnterConversation_statement is called when production conversation_statement is entered.
func (s *BaseTSqlParserListener) EnterConversation_statement(ctx *Conversation_statementContext) {}

// ExitConversation_statement is called when production conversation_statement is exited.
func (s *BaseTSqlParserListener) ExitConversation_statement(ctx *Conversation_statementContext) {}

// EnterMessage_statement is called when production message_statement is entered.
func (s *BaseTSqlParserListener) EnterMessage_statement(ctx *Message_statementContext) {}

// ExitMessage_statement is called when production message_statement is exited.
func (s *BaseTSqlParserListener) ExitMessage_statement(ctx *Message_statementContext) {}

// EnterMerge_statement is called when production merge_statement is entered.
func (s *BaseTSqlParserListener) EnterMerge_statement(ctx *Merge_statementContext) {}

// ExitMerge_statement is called when production merge_statement is exited.
func (s *BaseTSqlParserListener) ExitMerge_statement(ctx *Merge_statementContext) {}

// EnterWhen_matches is called when production when_matches is entered.
func (s *BaseTSqlParserListener) EnterWhen_matches(ctx *When_matchesContext) {}

// ExitWhen_matches is called when production when_matches is exited.
func (s *BaseTSqlParserListener) ExitWhen_matches(ctx *When_matchesContext) {}

// EnterMerge_matched is called when production merge_matched is entered.
func (s *BaseTSqlParserListener) EnterMerge_matched(ctx *Merge_matchedContext) {}

// ExitMerge_matched is called when production merge_matched is exited.
func (s *BaseTSqlParserListener) ExitMerge_matched(ctx *Merge_matchedContext) {}

// EnterMerge_not_matched is called when production merge_not_matched is entered.
func (s *BaseTSqlParserListener) EnterMerge_not_matched(ctx *Merge_not_matchedContext) {}

// ExitMerge_not_matched is called when production merge_not_matched is exited.
func (s *BaseTSqlParserListener) ExitMerge_not_matched(ctx *Merge_not_matchedContext) {}

// EnterDelete_statement is called when production delete_statement is entered.
func (s *BaseTSqlParserListener) EnterDelete_statement(ctx *Delete_statementContext) {}

// ExitDelete_statement is called when production delete_statement is exited.
func (s *BaseTSqlParserListener) ExitDelete_statement(ctx *Delete_statementContext) {}

// EnterDelete_statement_from is called when production delete_statement_from is entered.
func (s *BaseTSqlParserListener) EnterDelete_statement_from(ctx *Delete_statement_fromContext) {}

// ExitDelete_statement_from is called when production delete_statement_from is exited.
func (s *BaseTSqlParserListener) ExitDelete_statement_from(ctx *Delete_statement_fromContext) {}

// EnterInsert_statement is called when production insert_statement is entered.
func (s *BaseTSqlParserListener) EnterInsert_statement(ctx *Insert_statementContext) {}

// ExitInsert_statement is called when production insert_statement is exited.
func (s *BaseTSqlParserListener) ExitInsert_statement(ctx *Insert_statementContext) {}

// EnterInsert_statement_value is called when production insert_statement_value is entered.
func (s *BaseTSqlParserListener) EnterInsert_statement_value(ctx *Insert_statement_valueContext) {}

// ExitInsert_statement_value is called when production insert_statement_value is exited.
func (s *BaseTSqlParserListener) ExitInsert_statement_value(ctx *Insert_statement_valueContext) {}

// EnterReceive_statement is called when production receive_statement is entered.
func (s *BaseTSqlParserListener) EnterReceive_statement(ctx *Receive_statementContext) {}

// ExitReceive_statement is called when production receive_statement is exited.
func (s *BaseTSqlParserListener) ExitReceive_statement(ctx *Receive_statementContext) {}

// EnterSelect_statement_standalone is called when production select_statement_standalone is entered.
func (s *BaseTSqlParserListener) EnterSelect_statement_standalone(ctx *Select_statement_standaloneContext) {
}

// ExitSelect_statement_standalone is called when production select_statement_standalone is exited.
func (s *BaseTSqlParserListener) ExitSelect_statement_standalone(ctx *Select_statement_standaloneContext) {
}

// EnterSelect_statement is called when production select_statement is entered.
func (s *BaseTSqlParserListener) EnterSelect_statement(ctx *Select_statementContext) {}

// ExitSelect_statement is called when production select_statement is exited.
func (s *BaseTSqlParserListener) ExitSelect_statement(ctx *Select_statementContext) {}

// EnterTime is called when production time is entered.
func (s *BaseTSqlParserListener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *BaseTSqlParserListener) ExitTime(ctx *TimeContext) {}

// EnterUpdate_statement is called when production update_statement is entered.
func (s *BaseTSqlParserListener) EnterUpdate_statement(ctx *Update_statementContext) {}

// ExitUpdate_statement is called when production update_statement is exited.
func (s *BaseTSqlParserListener) ExitUpdate_statement(ctx *Update_statementContext) {}

// EnterOutput_clause is called when production output_clause is entered.
func (s *BaseTSqlParserListener) EnterOutput_clause(ctx *Output_clauseContext) {}

// ExitOutput_clause is called when production output_clause is exited.
func (s *BaseTSqlParserListener) ExitOutput_clause(ctx *Output_clauseContext) {}

// EnterOutput_dml_list_elem is called when production output_dml_list_elem is entered.
func (s *BaseTSqlParserListener) EnterOutput_dml_list_elem(ctx *Output_dml_list_elemContext) {}

// ExitOutput_dml_list_elem is called when production output_dml_list_elem is exited.
func (s *BaseTSqlParserListener) ExitOutput_dml_list_elem(ctx *Output_dml_list_elemContext) {}

// EnterCreate_database is called when production create_database is entered.
func (s *BaseTSqlParserListener) EnterCreate_database(ctx *Create_databaseContext) {}

// ExitCreate_database is called when production create_database is exited.
func (s *BaseTSqlParserListener) ExitCreate_database(ctx *Create_databaseContext) {}

// EnterCreate_index is called when production create_index is entered.
func (s *BaseTSqlParserListener) EnterCreate_index(ctx *Create_indexContext) {}

// ExitCreate_index is called when production create_index is exited.
func (s *BaseTSqlParserListener) ExitCreate_index(ctx *Create_indexContext) {}

// EnterAlter_index is called when production alter_index is entered.
func (s *BaseTSqlParserListener) EnterAlter_index(ctx *Alter_indexContext) {}

// ExitAlter_index is called when production alter_index is exited.
func (s *BaseTSqlParserListener) ExitAlter_index(ctx *Alter_indexContext) {}

// EnterRebuild_partition is called when production rebuild_partition is entered.
func (s *BaseTSqlParserListener) EnterRebuild_partition(ctx *Rebuild_partitionContext) {}

// ExitRebuild_partition is called when production rebuild_partition is exited.
func (s *BaseTSqlParserListener) ExitRebuild_partition(ctx *Rebuild_partitionContext) {}

// EnterCreate_columnstore_index is called when production create_columnstore_index is entered.
func (s *BaseTSqlParserListener) EnterCreate_columnstore_index(ctx *Create_columnstore_indexContext) {
}

// ExitCreate_columnstore_index is called when production create_columnstore_index is exited.
func (s *BaseTSqlParserListener) ExitCreate_columnstore_index(ctx *Create_columnstore_indexContext) {}

// EnterCreate_nonclustered_columnstore_index is called when production create_nonclustered_columnstore_index is entered.
func (s *BaseTSqlParserListener) EnterCreate_nonclustered_columnstore_index(ctx *Create_nonclustered_columnstore_indexContext) {
}

// ExitCreate_nonclustered_columnstore_index is called when production create_nonclustered_columnstore_index is exited.
func (s *BaseTSqlParserListener) ExitCreate_nonclustered_columnstore_index(ctx *Create_nonclustered_columnstore_indexContext) {
}

// EnterCreate_xml_index is called when production create_xml_index is entered.
func (s *BaseTSqlParserListener) EnterCreate_xml_index(ctx *Create_xml_indexContext) {}

// ExitCreate_xml_index is called when production create_xml_index is exited.
func (s *BaseTSqlParserListener) ExitCreate_xml_index(ctx *Create_xml_indexContext) {}

// EnterCreate_or_alter_procedure is called when production create_or_alter_procedure is entered.
func (s *BaseTSqlParserListener) EnterCreate_or_alter_procedure(ctx *Create_or_alter_procedureContext) {
}

// ExitCreate_or_alter_procedure is called when production create_or_alter_procedure is exited.
func (s *BaseTSqlParserListener) ExitCreate_or_alter_procedure(ctx *Create_or_alter_procedureContext) {
}

// EnterAs_external_name is called when production as_external_name is entered.
func (s *BaseTSqlParserListener) EnterAs_external_name(ctx *As_external_nameContext) {}

// ExitAs_external_name is called when production as_external_name is exited.
func (s *BaseTSqlParserListener) ExitAs_external_name(ctx *As_external_nameContext) {}

// EnterCreate_or_alter_trigger is called when production create_or_alter_trigger is entered.
func (s *BaseTSqlParserListener) EnterCreate_or_alter_trigger(ctx *Create_or_alter_triggerContext) {}

// ExitCreate_or_alter_trigger is called when production create_or_alter_trigger is exited.
func (s *BaseTSqlParserListener) ExitCreate_or_alter_trigger(ctx *Create_or_alter_triggerContext) {}

// EnterCreate_or_alter_dml_trigger is called when production create_or_alter_dml_trigger is entered.
func (s *BaseTSqlParserListener) EnterCreate_or_alter_dml_trigger(ctx *Create_or_alter_dml_triggerContext) {
}

// ExitCreate_or_alter_dml_trigger is called when production create_or_alter_dml_trigger is exited.
func (s *BaseTSqlParserListener) ExitCreate_or_alter_dml_trigger(ctx *Create_or_alter_dml_triggerContext) {
}

// EnterDml_trigger_option is called when production dml_trigger_option is entered.
func (s *BaseTSqlParserListener) EnterDml_trigger_option(ctx *Dml_trigger_optionContext) {}

// ExitDml_trigger_option is called when production dml_trigger_option is exited.
func (s *BaseTSqlParserListener) ExitDml_trigger_option(ctx *Dml_trigger_optionContext) {}

// EnterDml_trigger_operation is called when production dml_trigger_operation is entered.
func (s *BaseTSqlParserListener) EnterDml_trigger_operation(ctx *Dml_trigger_operationContext) {}

// ExitDml_trigger_operation is called when production dml_trigger_operation is exited.
func (s *BaseTSqlParserListener) ExitDml_trigger_operation(ctx *Dml_trigger_operationContext) {}

// EnterCreate_or_alter_ddl_trigger is called when production create_or_alter_ddl_trigger is entered.
func (s *BaseTSqlParserListener) EnterCreate_or_alter_ddl_trigger(ctx *Create_or_alter_ddl_triggerContext) {
}

// ExitCreate_or_alter_ddl_trigger is called when production create_or_alter_ddl_trigger is exited.
func (s *BaseTSqlParserListener) ExitCreate_or_alter_ddl_trigger(ctx *Create_or_alter_ddl_triggerContext) {
}

// EnterDdl_trigger_operation is called when production ddl_trigger_operation is entered.
func (s *BaseTSqlParserListener) EnterDdl_trigger_operation(ctx *Ddl_trigger_operationContext) {}

// ExitDdl_trigger_operation is called when production ddl_trigger_operation is exited.
func (s *BaseTSqlParserListener) ExitDdl_trigger_operation(ctx *Ddl_trigger_operationContext) {}

// EnterCreate_or_alter_function is called when production create_or_alter_function is entered.
func (s *BaseTSqlParserListener) EnterCreate_or_alter_function(ctx *Create_or_alter_functionContext) {
}

// ExitCreate_or_alter_function is called when production create_or_alter_function is exited.
func (s *BaseTSqlParserListener) ExitCreate_or_alter_function(ctx *Create_or_alter_functionContext) {}

// EnterFunc_body_returns_select is called when production func_body_returns_select is entered.
func (s *BaseTSqlParserListener) EnterFunc_body_returns_select(ctx *Func_body_returns_selectContext) {
}

// ExitFunc_body_returns_select is called when production func_body_returns_select is exited.
func (s *BaseTSqlParserListener) ExitFunc_body_returns_select(ctx *Func_body_returns_selectContext) {}

// EnterFunc_body_returns_table is called when production func_body_returns_table is entered.
func (s *BaseTSqlParserListener) EnterFunc_body_returns_table(ctx *Func_body_returns_tableContext) {}

// ExitFunc_body_returns_table is called when production func_body_returns_table is exited.
func (s *BaseTSqlParserListener) ExitFunc_body_returns_table(ctx *Func_body_returns_tableContext) {}

// EnterFunc_body_returns_scalar is called when production func_body_returns_scalar is entered.
func (s *BaseTSqlParserListener) EnterFunc_body_returns_scalar(ctx *Func_body_returns_scalarContext) {
}

// ExitFunc_body_returns_scalar is called when production func_body_returns_scalar is exited.
func (s *BaseTSqlParserListener) ExitFunc_body_returns_scalar(ctx *Func_body_returns_scalarContext) {}

// EnterProcedure_param is called when production procedure_param is entered.
func (s *BaseTSqlParserListener) EnterProcedure_param(ctx *Procedure_paramContext) {}

// ExitProcedure_param is called when production procedure_param is exited.
func (s *BaseTSqlParserListener) ExitProcedure_param(ctx *Procedure_paramContext) {}

// EnterProcedure_option is called when production procedure_option is entered.
func (s *BaseTSqlParserListener) EnterProcedure_option(ctx *Procedure_optionContext) {}

// ExitProcedure_option is called when production procedure_option is exited.
func (s *BaseTSqlParserListener) ExitProcedure_option(ctx *Procedure_optionContext) {}

// EnterFunction_option is called when production function_option is entered.
func (s *BaseTSqlParserListener) EnterFunction_option(ctx *Function_optionContext) {}

// ExitFunction_option is called when production function_option is exited.
func (s *BaseTSqlParserListener) ExitFunction_option(ctx *Function_optionContext) {}

// EnterCreate_statistics is called when production create_statistics is entered.
func (s *BaseTSqlParserListener) EnterCreate_statistics(ctx *Create_statisticsContext) {}

// ExitCreate_statistics is called when production create_statistics is exited.
func (s *BaseTSqlParserListener) ExitCreate_statistics(ctx *Create_statisticsContext) {}

// EnterUpdate_statistics is called when production update_statistics is entered.
func (s *BaseTSqlParserListener) EnterUpdate_statistics(ctx *Update_statisticsContext) {}

// ExitUpdate_statistics is called when production update_statistics is exited.
func (s *BaseTSqlParserListener) ExitUpdate_statistics(ctx *Update_statisticsContext) {}

// EnterCreate_table is called when production create_table is entered.
func (s *BaseTSqlParserListener) EnterCreate_table(ctx *Create_tableContext) {}

// ExitCreate_table is called when production create_table is exited.
func (s *BaseTSqlParserListener) ExitCreate_table(ctx *Create_tableContext) {}

// EnterTable_indices is called when production table_indices is entered.
func (s *BaseTSqlParserListener) EnterTable_indices(ctx *Table_indicesContext) {}

// ExitTable_indices is called when production table_indices is exited.
func (s *BaseTSqlParserListener) ExitTable_indices(ctx *Table_indicesContext) {}

// EnterTable_options is called when production table_options is entered.
func (s *BaseTSqlParserListener) EnterTable_options(ctx *Table_optionsContext) {}

// ExitTable_options is called when production table_options is exited.
func (s *BaseTSqlParserListener) ExitTable_options(ctx *Table_optionsContext) {}

// EnterCreate_view is called when production create_view is entered.
func (s *BaseTSqlParserListener) EnterCreate_view(ctx *Create_viewContext) {}

// ExitCreate_view is called when production create_view is exited.
func (s *BaseTSqlParserListener) ExitCreate_view(ctx *Create_viewContext) {}

// EnterView_attribute is called when production view_attribute is entered.
func (s *BaseTSqlParserListener) EnterView_attribute(ctx *View_attributeContext) {}

// ExitView_attribute is called when production view_attribute is exited.
func (s *BaseTSqlParserListener) ExitView_attribute(ctx *View_attributeContext) {}

// EnterAlter_table is called when production alter_table is entered.
func (s *BaseTSqlParserListener) EnterAlter_table(ctx *Alter_tableContext) {}

// ExitAlter_table is called when production alter_table is exited.
func (s *BaseTSqlParserListener) ExitAlter_table(ctx *Alter_tableContext) {}

// EnterSwitch_partition is called when production switch_partition is entered.
func (s *BaseTSqlParserListener) EnterSwitch_partition(ctx *Switch_partitionContext) {}

// ExitSwitch_partition is called when production switch_partition is exited.
func (s *BaseTSqlParserListener) ExitSwitch_partition(ctx *Switch_partitionContext) {}

// EnterLow_priority_lock_wait is called when production low_priority_lock_wait is entered.
func (s *BaseTSqlParserListener) EnterLow_priority_lock_wait(ctx *Low_priority_lock_waitContext) {}

// ExitLow_priority_lock_wait is called when production low_priority_lock_wait is exited.
func (s *BaseTSqlParserListener) ExitLow_priority_lock_wait(ctx *Low_priority_lock_waitContext) {}

// EnterAlter_database is called when production alter_database is entered.
func (s *BaseTSqlParserListener) EnterAlter_database(ctx *Alter_databaseContext) {}

// ExitAlter_database is called when production alter_database is exited.
func (s *BaseTSqlParserListener) ExitAlter_database(ctx *Alter_databaseContext) {}

// EnterAdd_or_modify_files is called when production add_or_modify_files is entered.
func (s *BaseTSqlParserListener) EnterAdd_or_modify_files(ctx *Add_or_modify_filesContext) {}

// ExitAdd_or_modify_files is called when production add_or_modify_files is exited.
func (s *BaseTSqlParserListener) ExitAdd_or_modify_files(ctx *Add_or_modify_filesContext) {}

// EnterFilespec is called when production filespec is entered.
func (s *BaseTSqlParserListener) EnterFilespec(ctx *FilespecContext) {}

// ExitFilespec is called when production filespec is exited.
func (s *BaseTSqlParserListener) ExitFilespec(ctx *FilespecContext) {}

// EnterAdd_or_modify_filegroups is called when production add_or_modify_filegroups is entered.
func (s *BaseTSqlParserListener) EnterAdd_or_modify_filegroups(ctx *Add_or_modify_filegroupsContext) {
}

// ExitAdd_or_modify_filegroups is called when production add_or_modify_filegroups is exited.
func (s *BaseTSqlParserListener) ExitAdd_or_modify_filegroups(ctx *Add_or_modify_filegroupsContext) {}

// EnterFilegroup_updatability_option is called when production filegroup_updatability_option is entered.
func (s *BaseTSqlParserListener) EnterFilegroup_updatability_option(ctx *Filegroup_updatability_optionContext) {
}

// ExitFilegroup_updatability_option is called when production filegroup_updatability_option is exited.
func (s *BaseTSqlParserListener) ExitFilegroup_updatability_option(ctx *Filegroup_updatability_optionContext) {
}

// EnterDatabase_optionspec is called when production database_optionspec is entered.
func (s *BaseTSqlParserListener) EnterDatabase_optionspec(ctx *Database_optionspecContext) {}

// ExitDatabase_optionspec is called when production database_optionspec is exited.
func (s *BaseTSqlParserListener) ExitDatabase_optionspec(ctx *Database_optionspecContext) {}

// EnterAuto_option is called when production auto_option is entered.
func (s *BaseTSqlParserListener) EnterAuto_option(ctx *Auto_optionContext) {}

// ExitAuto_option is called when production auto_option is exited.
func (s *BaseTSqlParserListener) ExitAuto_option(ctx *Auto_optionContext) {}

// EnterChange_tracking_option is called when production change_tracking_option is entered.
func (s *BaseTSqlParserListener) EnterChange_tracking_option(ctx *Change_tracking_optionContext) {}

// ExitChange_tracking_option is called when production change_tracking_option is exited.
func (s *BaseTSqlParserListener) ExitChange_tracking_option(ctx *Change_tracking_optionContext) {}

// EnterChange_tracking_option_list is called when production change_tracking_option_list is entered.
func (s *BaseTSqlParserListener) EnterChange_tracking_option_list(ctx *Change_tracking_option_listContext) {
}

// ExitChange_tracking_option_list is called when production change_tracking_option_list is exited.
func (s *BaseTSqlParserListener) ExitChange_tracking_option_list(ctx *Change_tracking_option_listContext) {
}

// EnterContainment_option is called when production containment_option is entered.
func (s *BaseTSqlParserListener) EnterContainment_option(ctx *Containment_optionContext) {}

// ExitContainment_option is called when production containment_option is exited.
func (s *BaseTSqlParserListener) ExitContainment_option(ctx *Containment_optionContext) {}

// EnterCursor_option is called when production cursor_option is entered.
func (s *BaseTSqlParserListener) EnterCursor_option(ctx *Cursor_optionContext) {}

// ExitCursor_option is called when production cursor_option is exited.
func (s *BaseTSqlParserListener) ExitCursor_option(ctx *Cursor_optionContext) {}

// EnterAlter_endpoint is called when production alter_endpoint is entered.
func (s *BaseTSqlParserListener) EnterAlter_endpoint(ctx *Alter_endpointContext) {}

// ExitAlter_endpoint is called when production alter_endpoint is exited.
func (s *BaseTSqlParserListener) ExitAlter_endpoint(ctx *Alter_endpointContext) {}

// EnterDatabase_mirroring_option is called when production database_mirroring_option is entered.
func (s *BaseTSqlParserListener) EnterDatabase_mirroring_option(ctx *Database_mirroring_optionContext) {
}

// ExitDatabase_mirroring_option is called when production database_mirroring_option is exited.
func (s *BaseTSqlParserListener) ExitDatabase_mirroring_option(ctx *Database_mirroring_optionContext) {
}

// EnterMirroring_set_option is called when production mirroring_set_option is entered.
func (s *BaseTSqlParserListener) EnterMirroring_set_option(ctx *Mirroring_set_optionContext) {}

// ExitMirroring_set_option is called when production mirroring_set_option is exited.
func (s *BaseTSqlParserListener) ExitMirroring_set_option(ctx *Mirroring_set_optionContext) {}

// EnterMirroring_partner is called when production mirroring_partner is entered.
func (s *BaseTSqlParserListener) EnterMirroring_partner(ctx *Mirroring_partnerContext) {}

// ExitMirroring_partner is called when production mirroring_partner is exited.
func (s *BaseTSqlParserListener) ExitMirroring_partner(ctx *Mirroring_partnerContext) {}

// EnterMirroring_witness is called when production mirroring_witness is entered.
func (s *BaseTSqlParserListener) EnterMirroring_witness(ctx *Mirroring_witnessContext) {}

// ExitMirroring_witness is called when production mirroring_witness is exited.
func (s *BaseTSqlParserListener) ExitMirroring_witness(ctx *Mirroring_witnessContext) {}

// EnterWitness_partner_equal is called when production witness_partner_equal is entered.
func (s *BaseTSqlParserListener) EnterWitness_partner_equal(ctx *Witness_partner_equalContext) {}

// ExitWitness_partner_equal is called when production witness_partner_equal is exited.
func (s *BaseTSqlParserListener) ExitWitness_partner_equal(ctx *Witness_partner_equalContext) {}

// EnterPartner_option is called when production partner_option is entered.
func (s *BaseTSqlParserListener) EnterPartner_option(ctx *Partner_optionContext) {}

// ExitPartner_option is called when production partner_option is exited.
func (s *BaseTSqlParserListener) ExitPartner_option(ctx *Partner_optionContext) {}

// EnterWitness_option is called when production witness_option is entered.
func (s *BaseTSqlParserListener) EnterWitness_option(ctx *Witness_optionContext) {}

// ExitWitness_option is called when production witness_option is exited.
func (s *BaseTSqlParserListener) ExitWitness_option(ctx *Witness_optionContext) {}

// EnterWitness_server is called when production witness_server is entered.
func (s *BaseTSqlParserListener) EnterWitness_server(ctx *Witness_serverContext) {}

// ExitWitness_server is called when production witness_server is exited.
func (s *BaseTSqlParserListener) ExitWitness_server(ctx *Witness_serverContext) {}

// EnterPartner_server is called when production partner_server is entered.
func (s *BaseTSqlParserListener) EnterPartner_server(ctx *Partner_serverContext) {}

// ExitPartner_server is called when production partner_server is exited.
func (s *BaseTSqlParserListener) ExitPartner_server(ctx *Partner_serverContext) {}

// EnterMirroring_host_port_seperator is called when production mirroring_host_port_seperator is entered.
func (s *BaseTSqlParserListener) EnterMirroring_host_port_seperator(ctx *Mirroring_host_port_seperatorContext) {
}

// ExitMirroring_host_port_seperator is called when production mirroring_host_port_seperator is exited.
func (s *BaseTSqlParserListener) ExitMirroring_host_port_seperator(ctx *Mirroring_host_port_seperatorContext) {
}

// EnterPartner_server_tcp_prefix is called when production partner_server_tcp_prefix is entered.
func (s *BaseTSqlParserListener) EnterPartner_server_tcp_prefix(ctx *Partner_server_tcp_prefixContext) {
}

// ExitPartner_server_tcp_prefix is called when production partner_server_tcp_prefix is exited.
func (s *BaseTSqlParserListener) ExitPartner_server_tcp_prefix(ctx *Partner_server_tcp_prefixContext) {
}

// EnterPort_number is called when production port_number is entered.
func (s *BaseTSqlParserListener) EnterPort_number(ctx *Port_numberContext) {}

// ExitPort_number is called when production port_number is exited.
func (s *BaseTSqlParserListener) ExitPort_number(ctx *Port_numberContext) {}

// EnterHost is called when production host is entered.
func (s *BaseTSqlParserListener) EnterHost(ctx *HostContext) {}

// ExitHost is called when production host is exited.
func (s *BaseTSqlParserListener) ExitHost(ctx *HostContext) {}

// EnterDate_correlation_optimization_option is called when production date_correlation_optimization_option is entered.
func (s *BaseTSqlParserListener) EnterDate_correlation_optimization_option(ctx *Date_correlation_optimization_optionContext) {
}

// ExitDate_correlation_optimization_option is called when production date_correlation_optimization_option is exited.
func (s *BaseTSqlParserListener) ExitDate_correlation_optimization_option(ctx *Date_correlation_optimization_optionContext) {
}

// EnterDb_encryption_option is called when production db_encryption_option is entered.
func (s *BaseTSqlParserListener) EnterDb_encryption_option(ctx *Db_encryption_optionContext) {}

// ExitDb_encryption_option is called when production db_encryption_option is exited.
func (s *BaseTSqlParserListener) ExitDb_encryption_option(ctx *Db_encryption_optionContext) {}

// EnterDb_state_option is called when production db_state_option is entered.
func (s *BaseTSqlParserListener) EnterDb_state_option(ctx *Db_state_optionContext) {}

// ExitDb_state_option is called when production db_state_option is exited.
func (s *BaseTSqlParserListener) ExitDb_state_option(ctx *Db_state_optionContext) {}

// EnterDb_update_option is called when production db_update_option is entered.
func (s *BaseTSqlParserListener) EnterDb_update_option(ctx *Db_update_optionContext) {}

// ExitDb_update_option is called when production db_update_option is exited.
func (s *BaseTSqlParserListener) ExitDb_update_option(ctx *Db_update_optionContext) {}

// EnterDb_user_access_option is called when production db_user_access_option is entered.
func (s *BaseTSqlParserListener) EnterDb_user_access_option(ctx *Db_user_access_optionContext) {}

// ExitDb_user_access_option is called when production db_user_access_option is exited.
func (s *BaseTSqlParserListener) ExitDb_user_access_option(ctx *Db_user_access_optionContext) {}

// EnterDelayed_durability_option is called when production delayed_durability_option is entered.
func (s *BaseTSqlParserListener) EnterDelayed_durability_option(ctx *Delayed_durability_optionContext) {
}

// ExitDelayed_durability_option is called when production delayed_durability_option is exited.
func (s *BaseTSqlParserListener) ExitDelayed_durability_option(ctx *Delayed_durability_optionContext) {
}

// EnterExternal_access_option is called when production external_access_option is entered.
func (s *BaseTSqlParserListener) EnterExternal_access_option(ctx *External_access_optionContext) {}

// ExitExternal_access_option is called when production external_access_option is exited.
func (s *BaseTSqlParserListener) ExitExternal_access_option(ctx *External_access_optionContext) {}

// EnterHadr_options is called when production hadr_options is entered.
func (s *BaseTSqlParserListener) EnterHadr_options(ctx *Hadr_optionsContext) {}

// ExitHadr_options is called when production hadr_options is exited.
func (s *BaseTSqlParserListener) ExitHadr_options(ctx *Hadr_optionsContext) {}

// EnterMixed_page_allocation_option is called when production mixed_page_allocation_option is entered.
func (s *BaseTSqlParserListener) EnterMixed_page_allocation_option(ctx *Mixed_page_allocation_optionContext) {
}

// ExitMixed_page_allocation_option is called when production mixed_page_allocation_option is exited.
func (s *BaseTSqlParserListener) ExitMixed_page_allocation_option(ctx *Mixed_page_allocation_optionContext) {
}

// EnterParameterization_option is called when production parameterization_option is entered.
func (s *BaseTSqlParserListener) EnterParameterization_option(ctx *Parameterization_optionContext) {}

// ExitParameterization_option is called when production parameterization_option is exited.
func (s *BaseTSqlParserListener) ExitParameterization_option(ctx *Parameterization_optionContext) {}

// EnterRecovery_option is called when production recovery_option is entered.
func (s *BaseTSqlParserListener) EnterRecovery_option(ctx *Recovery_optionContext) {}

// ExitRecovery_option is called when production recovery_option is exited.
func (s *BaseTSqlParserListener) ExitRecovery_option(ctx *Recovery_optionContext) {}

// EnterService_broker_option is called when production service_broker_option is entered.
func (s *BaseTSqlParserListener) EnterService_broker_option(ctx *Service_broker_optionContext) {}

// ExitService_broker_option is called when production service_broker_option is exited.
func (s *BaseTSqlParserListener) ExitService_broker_option(ctx *Service_broker_optionContext) {}

// EnterSnapshot_option is called when production snapshot_option is entered.
func (s *BaseTSqlParserListener) EnterSnapshot_option(ctx *Snapshot_optionContext) {}

// ExitSnapshot_option is called when production snapshot_option is exited.
func (s *BaseTSqlParserListener) ExitSnapshot_option(ctx *Snapshot_optionContext) {}

// EnterSql_option is called when production sql_option is entered.
func (s *BaseTSqlParserListener) EnterSql_option(ctx *Sql_optionContext) {}

// ExitSql_option is called when production sql_option is exited.
func (s *BaseTSqlParserListener) ExitSql_option(ctx *Sql_optionContext) {}

// EnterTarget_recovery_time_option is called when production target_recovery_time_option is entered.
func (s *BaseTSqlParserListener) EnterTarget_recovery_time_option(ctx *Target_recovery_time_optionContext) {
}

// ExitTarget_recovery_time_option is called when production target_recovery_time_option is exited.
func (s *BaseTSqlParserListener) ExitTarget_recovery_time_option(ctx *Target_recovery_time_optionContext) {
}

// EnterTermination is called when production termination is entered.
func (s *BaseTSqlParserListener) EnterTermination(ctx *TerminationContext) {}

// ExitTermination is called when production termination is exited.
func (s *BaseTSqlParserListener) ExitTermination(ctx *TerminationContext) {}

// EnterDrop_index is called when production drop_index is entered.
func (s *BaseTSqlParserListener) EnterDrop_index(ctx *Drop_indexContext) {}

// ExitDrop_index is called when production drop_index is exited.
func (s *BaseTSqlParserListener) ExitDrop_index(ctx *Drop_indexContext) {}

// EnterDrop_relational_or_xml_or_spatial_index is called when production drop_relational_or_xml_or_spatial_index is entered.
func (s *BaseTSqlParserListener) EnterDrop_relational_or_xml_or_spatial_index(ctx *Drop_relational_or_xml_or_spatial_indexContext) {
}

// ExitDrop_relational_or_xml_or_spatial_index is called when production drop_relational_or_xml_or_spatial_index is exited.
func (s *BaseTSqlParserListener) ExitDrop_relational_or_xml_or_spatial_index(ctx *Drop_relational_or_xml_or_spatial_indexContext) {
}

// EnterDrop_backward_compatible_index is called when production drop_backward_compatible_index is entered.
func (s *BaseTSqlParserListener) EnterDrop_backward_compatible_index(ctx *Drop_backward_compatible_indexContext) {
}

// ExitDrop_backward_compatible_index is called when production drop_backward_compatible_index is exited.
func (s *BaseTSqlParserListener) ExitDrop_backward_compatible_index(ctx *Drop_backward_compatible_indexContext) {
}

// EnterDrop_procedure is called when production drop_procedure is entered.
func (s *BaseTSqlParserListener) EnterDrop_procedure(ctx *Drop_procedureContext) {}

// ExitDrop_procedure is called when production drop_procedure is exited.
func (s *BaseTSqlParserListener) ExitDrop_procedure(ctx *Drop_procedureContext) {}

// EnterDrop_trigger is called when production drop_trigger is entered.
func (s *BaseTSqlParserListener) EnterDrop_trigger(ctx *Drop_triggerContext) {}

// ExitDrop_trigger is called when production drop_trigger is exited.
func (s *BaseTSqlParserListener) ExitDrop_trigger(ctx *Drop_triggerContext) {}

// EnterDrop_dml_trigger is called when production drop_dml_trigger is entered.
func (s *BaseTSqlParserListener) EnterDrop_dml_trigger(ctx *Drop_dml_triggerContext) {}

// ExitDrop_dml_trigger is called when production drop_dml_trigger is exited.
func (s *BaseTSqlParserListener) ExitDrop_dml_trigger(ctx *Drop_dml_triggerContext) {}

// EnterDrop_ddl_trigger is called when production drop_ddl_trigger is entered.
func (s *BaseTSqlParserListener) EnterDrop_ddl_trigger(ctx *Drop_ddl_triggerContext) {}

// ExitDrop_ddl_trigger is called when production drop_ddl_trigger is exited.
func (s *BaseTSqlParserListener) ExitDrop_ddl_trigger(ctx *Drop_ddl_triggerContext) {}

// EnterDrop_function is called when production drop_function is entered.
func (s *BaseTSqlParserListener) EnterDrop_function(ctx *Drop_functionContext) {}

// ExitDrop_function is called when production drop_function is exited.
func (s *BaseTSqlParserListener) ExitDrop_function(ctx *Drop_functionContext) {}

// EnterDrop_statistics is called when production drop_statistics is entered.
func (s *BaseTSqlParserListener) EnterDrop_statistics(ctx *Drop_statisticsContext) {}

// ExitDrop_statistics is called when production drop_statistics is exited.
func (s *BaseTSqlParserListener) ExitDrop_statistics(ctx *Drop_statisticsContext) {}

// EnterDrop_table is called when production drop_table is entered.
func (s *BaseTSqlParserListener) EnterDrop_table(ctx *Drop_tableContext) {}

// ExitDrop_table is called when production drop_table is exited.
func (s *BaseTSqlParserListener) ExitDrop_table(ctx *Drop_tableContext) {}

// EnterDrop_view is called when production drop_view is entered.
func (s *BaseTSqlParserListener) EnterDrop_view(ctx *Drop_viewContext) {}

// ExitDrop_view is called when production drop_view is exited.
func (s *BaseTSqlParserListener) ExitDrop_view(ctx *Drop_viewContext) {}

// EnterCreate_type is called when production create_type is entered.
func (s *BaseTSqlParserListener) EnterCreate_type(ctx *Create_typeContext) {}

// ExitCreate_type is called when production create_type is exited.
func (s *BaseTSqlParserListener) ExitCreate_type(ctx *Create_typeContext) {}

// EnterDrop_type is called when production drop_type is entered.
func (s *BaseTSqlParserListener) EnterDrop_type(ctx *Drop_typeContext) {}

// ExitDrop_type is called when production drop_type is exited.
func (s *BaseTSqlParserListener) ExitDrop_type(ctx *Drop_typeContext) {}

// EnterRowset_function_limited is called when production rowset_function_limited is entered.
func (s *BaseTSqlParserListener) EnterRowset_function_limited(ctx *Rowset_function_limitedContext) {}

// ExitRowset_function_limited is called when production rowset_function_limited is exited.
func (s *BaseTSqlParserListener) ExitRowset_function_limited(ctx *Rowset_function_limitedContext) {}

// EnterOpenquery is called when production openquery is entered.
func (s *BaseTSqlParserListener) EnterOpenquery(ctx *OpenqueryContext) {}

// ExitOpenquery is called when production openquery is exited.
func (s *BaseTSqlParserListener) ExitOpenquery(ctx *OpenqueryContext) {}

// EnterOpendatasource is called when production opendatasource is entered.
func (s *BaseTSqlParserListener) EnterOpendatasource(ctx *OpendatasourceContext) {}

// ExitOpendatasource is called when production opendatasource is exited.
func (s *BaseTSqlParserListener) ExitOpendatasource(ctx *OpendatasourceContext) {}

// EnterDeclare_statement is called when production declare_statement is entered.
func (s *BaseTSqlParserListener) EnterDeclare_statement(ctx *Declare_statementContext) {}

// ExitDeclare_statement is called when production declare_statement is exited.
func (s *BaseTSqlParserListener) ExitDeclare_statement(ctx *Declare_statementContext) {}

// EnterXml_declaration is called when production xml_declaration is entered.
func (s *BaseTSqlParserListener) EnterXml_declaration(ctx *Xml_declarationContext) {}

// ExitXml_declaration is called when production xml_declaration is exited.
func (s *BaseTSqlParserListener) ExitXml_declaration(ctx *Xml_declarationContext) {}

// EnterCursor_statement is called when production cursor_statement is entered.
func (s *BaseTSqlParserListener) EnterCursor_statement(ctx *Cursor_statementContext) {}

// ExitCursor_statement is called when production cursor_statement is exited.
func (s *BaseTSqlParserListener) ExitCursor_statement(ctx *Cursor_statementContext) {}

// EnterBackup_database is called when production backup_database is entered.
func (s *BaseTSqlParserListener) EnterBackup_database(ctx *Backup_databaseContext) {}

// ExitBackup_database is called when production backup_database is exited.
func (s *BaseTSqlParserListener) ExitBackup_database(ctx *Backup_databaseContext) {}

// EnterBackup_log is called when production backup_log is entered.
func (s *BaseTSqlParserListener) EnterBackup_log(ctx *Backup_logContext) {}

// ExitBackup_log is called when production backup_log is exited.
func (s *BaseTSqlParserListener) ExitBackup_log(ctx *Backup_logContext) {}

// EnterBackup_certificate is called when production backup_certificate is entered.
func (s *BaseTSqlParserListener) EnterBackup_certificate(ctx *Backup_certificateContext) {}

// ExitBackup_certificate is called when production backup_certificate is exited.
func (s *BaseTSqlParserListener) ExitBackup_certificate(ctx *Backup_certificateContext) {}

// EnterBackup_master_key is called when production backup_master_key is entered.
func (s *BaseTSqlParserListener) EnterBackup_master_key(ctx *Backup_master_keyContext) {}

// ExitBackup_master_key is called when production backup_master_key is exited.
func (s *BaseTSqlParserListener) ExitBackup_master_key(ctx *Backup_master_keyContext) {}

// EnterBackup_service_master_key is called when production backup_service_master_key is entered.
func (s *BaseTSqlParserListener) EnterBackup_service_master_key(ctx *Backup_service_master_keyContext) {
}

// ExitBackup_service_master_key is called when production backup_service_master_key is exited.
func (s *BaseTSqlParserListener) ExitBackup_service_master_key(ctx *Backup_service_master_keyContext) {
}

// EnterKill_statement is called when production kill_statement is entered.
func (s *BaseTSqlParserListener) EnterKill_statement(ctx *Kill_statementContext) {}

// ExitKill_statement is called when production kill_statement is exited.
func (s *BaseTSqlParserListener) ExitKill_statement(ctx *Kill_statementContext) {}

// EnterKill_process is called when production kill_process is entered.
func (s *BaseTSqlParserListener) EnterKill_process(ctx *Kill_processContext) {}

// ExitKill_process is called when production kill_process is exited.
func (s *BaseTSqlParserListener) ExitKill_process(ctx *Kill_processContext) {}

// EnterKill_query_notification is called when production kill_query_notification is entered.
func (s *BaseTSqlParserListener) EnterKill_query_notification(ctx *Kill_query_notificationContext) {}

// ExitKill_query_notification is called when production kill_query_notification is exited.
func (s *BaseTSqlParserListener) ExitKill_query_notification(ctx *Kill_query_notificationContext) {}

// EnterKill_stats_job is called when production kill_stats_job is entered.
func (s *BaseTSqlParserListener) EnterKill_stats_job(ctx *Kill_stats_jobContext) {}

// ExitKill_stats_job is called when production kill_stats_job is exited.
func (s *BaseTSqlParserListener) ExitKill_stats_job(ctx *Kill_stats_jobContext) {}

// EnterExecute_statement is called when production execute_statement is entered.
func (s *BaseTSqlParserListener) EnterExecute_statement(ctx *Execute_statementContext) {}

// ExitExecute_statement is called when production execute_statement is exited.
func (s *BaseTSqlParserListener) ExitExecute_statement(ctx *Execute_statementContext) {}

// EnterExecute_body_batch is called when production execute_body_batch is entered.
func (s *BaseTSqlParserListener) EnterExecute_body_batch(ctx *Execute_body_batchContext) {}

// ExitExecute_body_batch is called when production execute_body_batch is exited.
func (s *BaseTSqlParserListener) ExitExecute_body_batch(ctx *Execute_body_batchContext) {}

// EnterExecute_body is called when production execute_body is entered.
func (s *BaseTSqlParserListener) EnterExecute_body(ctx *Execute_bodyContext) {}

// ExitExecute_body is called when production execute_body is exited.
func (s *BaseTSqlParserListener) ExitExecute_body(ctx *Execute_bodyContext) {}

// EnterExecute_statement_arg is called when production execute_statement_arg is entered.
func (s *BaseTSqlParserListener) EnterExecute_statement_arg(ctx *Execute_statement_argContext) {}

// ExitExecute_statement_arg is called when production execute_statement_arg is exited.
func (s *BaseTSqlParserListener) ExitExecute_statement_arg(ctx *Execute_statement_argContext) {}

// EnterExecute_statement_arg_named is called when production execute_statement_arg_named is entered.
func (s *BaseTSqlParserListener) EnterExecute_statement_arg_named(ctx *Execute_statement_arg_namedContext) {
}

// ExitExecute_statement_arg_named is called when production execute_statement_arg_named is exited.
func (s *BaseTSqlParserListener) ExitExecute_statement_arg_named(ctx *Execute_statement_arg_namedContext) {
}

// EnterExecute_statement_arg_unnamed is called when production execute_statement_arg_unnamed is entered.
func (s *BaseTSqlParserListener) EnterExecute_statement_arg_unnamed(ctx *Execute_statement_arg_unnamedContext) {
}

// ExitExecute_statement_arg_unnamed is called when production execute_statement_arg_unnamed is exited.
func (s *BaseTSqlParserListener) ExitExecute_statement_arg_unnamed(ctx *Execute_statement_arg_unnamedContext) {
}

// EnterExecute_parameter is called when production execute_parameter is entered.
func (s *BaseTSqlParserListener) EnterExecute_parameter(ctx *Execute_parameterContext) {}

// ExitExecute_parameter is called when production execute_parameter is exited.
func (s *BaseTSqlParserListener) ExitExecute_parameter(ctx *Execute_parameterContext) {}

// EnterExecute_var_string is called when production execute_var_string is entered.
func (s *BaseTSqlParserListener) EnterExecute_var_string(ctx *Execute_var_stringContext) {}

// ExitExecute_var_string is called when production execute_var_string is exited.
func (s *BaseTSqlParserListener) ExitExecute_var_string(ctx *Execute_var_stringContext) {}

// EnterSecurity_statement is called when production security_statement is entered.
func (s *BaseTSqlParserListener) EnterSecurity_statement(ctx *Security_statementContext) {}

// ExitSecurity_statement is called when production security_statement is exited.
func (s *BaseTSqlParserListener) ExitSecurity_statement(ctx *Security_statementContext) {}

// EnterPrincipal_id is called when production principal_id is entered.
func (s *BaseTSqlParserListener) EnterPrincipal_id(ctx *Principal_idContext) {}

// ExitPrincipal_id is called when production principal_id is exited.
func (s *BaseTSqlParserListener) ExitPrincipal_id(ctx *Principal_idContext) {}

// EnterCreate_certificate is called when production create_certificate is entered.
func (s *BaseTSqlParserListener) EnterCreate_certificate(ctx *Create_certificateContext) {}

// ExitCreate_certificate is called when production create_certificate is exited.
func (s *BaseTSqlParserListener) ExitCreate_certificate(ctx *Create_certificateContext) {}

// EnterExisting_keys is called when production existing_keys is entered.
func (s *BaseTSqlParserListener) EnterExisting_keys(ctx *Existing_keysContext) {}

// ExitExisting_keys is called when production existing_keys is exited.
func (s *BaseTSqlParserListener) ExitExisting_keys(ctx *Existing_keysContext) {}

// EnterPrivate_key_options is called when production private_key_options is entered.
func (s *BaseTSqlParserListener) EnterPrivate_key_options(ctx *Private_key_optionsContext) {}

// ExitPrivate_key_options is called when production private_key_options is exited.
func (s *BaseTSqlParserListener) ExitPrivate_key_options(ctx *Private_key_optionsContext) {}

// EnterGenerate_new_keys is called when production generate_new_keys is entered.
func (s *BaseTSqlParserListener) EnterGenerate_new_keys(ctx *Generate_new_keysContext) {}

// ExitGenerate_new_keys is called when production generate_new_keys is exited.
func (s *BaseTSqlParserListener) ExitGenerate_new_keys(ctx *Generate_new_keysContext) {}

// EnterDate_options is called when production date_options is entered.
func (s *BaseTSqlParserListener) EnterDate_options(ctx *Date_optionsContext) {}

// ExitDate_options is called when production date_options is exited.
func (s *BaseTSqlParserListener) ExitDate_options(ctx *Date_optionsContext) {}

// EnterOpen_key is called when production open_key is entered.
func (s *BaseTSqlParserListener) EnterOpen_key(ctx *Open_keyContext) {}

// ExitOpen_key is called when production open_key is exited.
func (s *BaseTSqlParserListener) ExitOpen_key(ctx *Open_keyContext) {}

// EnterClose_key is called when production close_key is entered.
func (s *BaseTSqlParserListener) EnterClose_key(ctx *Close_keyContext) {}

// ExitClose_key is called when production close_key is exited.
func (s *BaseTSqlParserListener) ExitClose_key(ctx *Close_keyContext) {}

// EnterCreate_key is called when production create_key is entered.
func (s *BaseTSqlParserListener) EnterCreate_key(ctx *Create_keyContext) {}

// ExitCreate_key is called when production create_key is exited.
func (s *BaseTSqlParserListener) ExitCreate_key(ctx *Create_keyContext) {}

// EnterKey_options is called when production key_options is entered.
func (s *BaseTSqlParserListener) EnterKey_options(ctx *Key_optionsContext) {}

// ExitKey_options is called when production key_options is exited.
func (s *BaseTSqlParserListener) ExitKey_options(ctx *Key_optionsContext) {}

// EnterAlgorithm is called when production algorithm is entered.
func (s *BaseTSqlParserListener) EnterAlgorithm(ctx *AlgorithmContext) {}

// ExitAlgorithm is called when production algorithm is exited.
func (s *BaseTSqlParserListener) ExitAlgorithm(ctx *AlgorithmContext) {}

// EnterEncryption_mechanism is called when production encryption_mechanism is entered.
func (s *BaseTSqlParserListener) EnterEncryption_mechanism(ctx *Encryption_mechanismContext) {}

// ExitEncryption_mechanism is called when production encryption_mechanism is exited.
func (s *BaseTSqlParserListener) ExitEncryption_mechanism(ctx *Encryption_mechanismContext) {}

// EnterDecryption_mechanism is called when production decryption_mechanism is entered.
func (s *BaseTSqlParserListener) EnterDecryption_mechanism(ctx *Decryption_mechanismContext) {}

// ExitDecryption_mechanism is called when production decryption_mechanism is exited.
func (s *BaseTSqlParserListener) ExitDecryption_mechanism(ctx *Decryption_mechanismContext) {}

// EnterGrant_permission is called when production grant_permission is entered.
func (s *BaseTSqlParserListener) EnterGrant_permission(ctx *Grant_permissionContext) {}

// ExitGrant_permission is called when production grant_permission is exited.
func (s *BaseTSqlParserListener) ExitGrant_permission(ctx *Grant_permissionContext) {}

// EnterSet_statement is called when production set_statement is entered.
func (s *BaseTSqlParserListener) EnterSet_statement(ctx *Set_statementContext) {}

// ExitSet_statement is called when production set_statement is exited.
func (s *BaseTSqlParserListener) ExitSet_statement(ctx *Set_statementContext) {}

// EnterTransaction_statement is called when production transaction_statement is entered.
func (s *BaseTSqlParserListener) EnterTransaction_statement(ctx *Transaction_statementContext) {}

// ExitTransaction_statement is called when production transaction_statement is exited.
func (s *BaseTSqlParserListener) ExitTransaction_statement(ctx *Transaction_statementContext) {}

// EnterGo_batch_statement is called when production go_batch_statement is entered.
func (s *BaseTSqlParserListener) EnterGo_batch_statement(ctx *Go_batch_statementContext) {}

// ExitGo_batch_statement is called when production go_batch_statement is exited.
func (s *BaseTSqlParserListener) ExitGo_batch_statement(ctx *Go_batch_statementContext) {}

// EnterGo_statement is called when production go_statement is entered.
func (s *BaseTSqlParserListener) EnterGo_statement(ctx *Go_statementContext) {}

// ExitGo_statement is called when production go_statement is exited.
func (s *BaseTSqlParserListener) ExitGo_statement(ctx *Go_statementContext) {}

// EnterUse_statement is called when production use_statement is entered.
func (s *BaseTSqlParserListener) EnterUse_statement(ctx *Use_statementContext) {}

// ExitUse_statement is called when production use_statement is exited.
func (s *BaseTSqlParserListener) ExitUse_statement(ctx *Use_statementContext) {}

// EnterSetuser_statement is called when production setuser_statement is entered.
func (s *BaseTSqlParserListener) EnterSetuser_statement(ctx *Setuser_statementContext) {}

// ExitSetuser_statement is called when production setuser_statement is exited.
func (s *BaseTSqlParserListener) ExitSetuser_statement(ctx *Setuser_statementContext) {}

// EnterReconfigure_statement is called when production reconfigure_statement is entered.
func (s *BaseTSqlParserListener) EnterReconfigure_statement(ctx *Reconfigure_statementContext) {}

// ExitReconfigure_statement is called when production reconfigure_statement is exited.
func (s *BaseTSqlParserListener) ExitReconfigure_statement(ctx *Reconfigure_statementContext) {}

// EnterShutdown_statement is called when production shutdown_statement is entered.
func (s *BaseTSqlParserListener) EnterShutdown_statement(ctx *Shutdown_statementContext) {}

// ExitShutdown_statement is called when production shutdown_statement is exited.
func (s *BaseTSqlParserListener) ExitShutdown_statement(ctx *Shutdown_statementContext) {}

// EnterCheckpoint_statement is called when production checkpoint_statement is entered.
func (s *BaseTSqlParserListener) EnterCheckpoint_statement(ctx *Checkpoint_statementContext) {}

// ExitCheckpoint_statement is called when production checkpoint_statement is exited.
func (s *BaseTSqlParserListener) ExitCheckpoint_statement(ctx *Checkpoint_statementContext) {}

// EnterDbcc_special is called when production dbcc_special is entered.
func (s *BaseTSqlParserListener) EnterDbcc_special(ctx *Dbcc_specialContext) {}

// ExitDbcc_special is called when production dbcc_special is exited.
func (s *BaseTSqlParserListener) ExitDbcc_special(ctx *Dbcc_specialContext) {}

// EnterDbcc_clause is called when production dbcc_clause is entered.
func (s *BaseTSqlParserListener) EnterDbcc_clause(ctx *Dbcc_clauseContext) {}

// ExitDbcc_clause is called when production dbcc_clause is exited.
func (s *BaseTSqlParserListener) ExitDbcc_clause(ctx *Dbcc_clauseContext) {}

// EnterDbcc_command is called when production dbcc_command is entered.
func (s *BaseTSqlParserListener) EnterDbcc_command(ctx *Dbcc_commandContext) {}

// ExitDbcc_command is called when production dbcc_command is exited.
func (s *BaseTSqlParserListener) ExitDbcc_command(ctx *Dbcc_commandContext) {}

// EnterDbcc_options is called when production dbcc_options is entered.
func (s *BaseTSqlParserListener) EnterDbcc_options(ctx *Dbcc_optionsContext) {}

// ExitDbcc_options is called when production dbcc_options is exited.
func (s *BaseTSqlParserListener) ExitDbcc_options(ctx *Dbcc_optionsContext) {}

// EnterExecute_clause is called when production execute_clause is entered.
func (s *BaseTSqlParserListener) EnterExecute_clause(ctx *Execute_clauseContext) {}

// ExitExecute_clause is called when production execute_clause is exited.
func (s *BaseTSqlParserListener) ExitExecute_clause(ctx *Execute_clauseContext) {}

// EnterDeclare_local is called when production declare_local is entered.
func (s *BaseTSqlParserListener) EnterDeclare_local(ctx *Declare_localContext) {}

// ExitDeclare_local is called when production declare_local is exited.
func (s *BaseTSqlParserListener) ExitDeclare_local(ctx *Declare_localContext) {}

// EnterTable_type_definition is called when production table_type_definition is entered.
func (s *BaseTSqlParserListener) EnterTable_type_definition(ctx *Table_type_definitionContext) {}

// ExitTable_type_definition is called when production table_type_definition is exited.
func (s *BaseTSqlParserListener) ExitTable_type_definition(ctx *Table_type_definitionContext) {}

// EnterTable_type_indices is called when production table_type_indices is entered.
func (s *BaseTSqlParserListener) EnterTable_type_indices(ctx *Table_type_indicesContext) {}

// ExitTable_type_indices is called when production table_type_indices is exited.
func (s *BaseTSqlParserListener) ExitTable_type_indices(ctx *Table_type_indicesContext) {}

// EnterXml_type_definition is called when production xml_type_definition is entered.
func (s *BaseTSqlParserListener) EnterXml_type_definition(ctx *Xml_type_definitionContext) {}

// ExitXml_type_definition is called when production xml_type_definition is exited.
func (s *BaseTSqlParserListener) ExitXml_type_definition(ctx *Xml_type_definitionContext) {}

// EnterXml_schema_collection is called when production xml_schema_collection is entered.
func (s *BaseTSqlParserListener) EnterXml_schema_collection(ctx *Xml_schema_collectionContext) {}

// ExitXml_schema_collection is called when production xml_schema_collection is exited.
func (s *BaseTSqlParserListener) ExitXml_schema_collection(ctx *Xml_schema_collectionContext) {}

// EnterColumn_def_table_constraints is called when production column_def_table_constraints is entered.
func (s *BaseTSqlParserListener) EnterColumn_def_table_constraints(ctx *Column_def_table_constraintsContext) {
}

// ExitColumn_def_table_constraints is called when production column_def_table_constraints is exited.
func (s *BaseTSqlParserListener) ExitColumn_def_table_constraints(ctx *Column_def_table_constraintsContext) {
}

// EnterColumn_def_table_constraint is called when production column_def_table_constraint is entered.
func (s *BaseTSqlParserListener) EnterColumn_def_table_constraint(ctx *Column_def_table_constraintContext) {
}

// ExitColumn_def_table_constraint is called when production column_def_table_constraint is exited.
func (s *BaseTSqlParserListener) ExitColumn_def_table_constraint(ctx *Column_def_table_constraintContext) {
}

// EnterColumn_definition is called when production column_definition is entered.
func (s *BaseTSqlParserListener) EnterColumn_definition(ctx *Column_definitionContext) {}

// ExitColumn_definition is called when production column_definition is exited.
func (s *BaseTSqlParserListener) ExitColumn_definition(ctx *Column_definitionContext) {}

// EnterMaterialized_column_definition is called when production materialized_column_definition is entered.
func (s *BaseTSqlParserListener) EnterMaterialized_column_definition(ctx *Materialized_column_definitionContext) {
}

// ExitMaterialized_column_definition is called when production materialized_column_definition is exited.
func (s *BaseTSqlParserListener) ExitMaterialized_column_definition(ctx *Materialized_column_definitionContext) {
}

// EnterColumn_constraint is called when production column_constraint is entered.
func (s *BaseTSqlParserListener) EnterColumn_constraint(ctx *Column_constraintContext) {}

// ExitColumn_constraint is called when production column_constraint is exited.
func (s *BaseTSqlParserListener) ExitColumn_constraint(ctx *Column_constraintContext) {}

// EnterTable_constraint is called when production table_constraint is entered.
func (s *BaseTSqlParserListener) EnterTable_constraint(ctx *Table_constraintContext) {}

// ExitTable_constraint is called when production table_constraint is exited.
func (s *BaseTSqlParserListener) ExitTable_constraint(ctx *Table_constraintContext) {}

// EnterOn_delete is called when production on_delete is entered.
func (s *BaseTSqlParserListener) EnterOn_delete(ctx *On_deleteContext) {}

// ExitOn_delete is called when production on_delete is exited.
func (s *BaseTSqlParserListener) ExitOn_delete(ctx *On_deleteContext) {}

// EnterOn_update is called when production on_update is entered.
func (s *BaseTSqlParserListener) EnterOn_update(ctx *On_updateContext) {}

// ExitOn_update is called when production on_update is exited.
func (s *BaseTSqlParserListener) ExitOn_update(ctx *On_updateContext) {}

// EnterIndex_options is called when production index_options is entered.
func (s *BaseTSqlParserListener) EnterIndex_options(ctx *Index_optionsContext) {}

// ExitIndex_options is called when production index_options is exited.
func (s *BaseTSqlParserListener) ExitIndex_options(ctx *Index_optionsContext) {}

// EnterIndex_option is called when production index_option is entered.
func (s *BaseTSqlParserListener) EnterIndex_option(ctx *Index_optionContext) {}

// ExitIndex_option is called when production index_option is exited.
func (s *BaseTSqlParserListener) ExitIndex_option(ctx *Index_optionContext) {}

// EnterDeclare_cursor is called when production declare_cursor is entered.
func (s *BaseTSqlParserListener) EnterDeclare_cursor(ctx *Declare_cursorContext) {}

// ExitDeclare_cursor is called when production declare_cursor is exited.
func (s *BaseTSqlParserListener) ExitDeclare_cursor(ctx *Declare_cursorContext) {}

// EnterDeclare_set_cursor_common is called when production declare_set_cursor_common is entered.
func (s *BaseTSqlParserListener) EnterDeclare_set_cursor_common(ctx *Declare_set_cursor_commonContext) {
}

// ExitDeclare_set_cursor_common is called when production declare_set_cursor_common is exited.
func (s *BaseTSqlParserListener) ExitDeclare_set_cursor_common(ctx *Declare_set_cursor_commonContext) {
}

// EnterDeclare_set_cursor_common_partial is called when production declare_set_cursor_common_partial is entered.
func (s *BaseTSqlParserListener) EnterDeclare_set_cursor_common_partial(ctx *Declare_set_cursor_common_partialContext) {
}

// ExitDeclare_set_cursor_common_partial is called when production declare_set_cursor_common_partial is exited.
func (s *BaseTSqlParserListener) ExitDeclare_set_cursor_common_partial(ctx *Declare_set_cursor_common_partialContext) {
}

// EnterFetch_cursor is called when production fetch_cursor is entered.
func (s *BaseTSqlParserListener) EnterFetch_cursor(ctx *Fetch_cursorContext) {}

// ExitFetch_cursor is called when production fetch_cursor is exited.
func (s *BaseTSqlParserListener) ExitFetch_cursor(ctx *Fetch_cursorContext) {}

// EnterSet_special is called when production set_special is entered.
func (s *BaseTSqlParserListener) EnterSet_special(ctx *Set_specialContext) {}

// ExitSet_special is called when production set_special is exited.
func (s *BaseTSqlParserListener) ExitSet_special(ctx *Set_specialContext) {}

// EnterSpecial_list is called when production special_list is entered.
func (s *BaseTSqlParserListener) EnterSpecial_list(ctx *Special_listContext) {}

// ExitSpecial_list is called when production special_list is exited.
func (s *BaseTSqlParserListener) ExitSpecial_list(ctx *Special_listContext) {}

// EnterConstant_LOCAL_ID is called when production constant_LOCAL_ID is entered.
func (s *BaseTSqlParserListener) EnterConstant_LOCAL_ID(ctx *Constant_LOCAL_IDContext) {}

// ExitConstant_LOCAL_ID is called when production constant_LOCAL_ID is exited.
func (s *BaseTSqlParserListener) ExitConstant_LOCAL_ID(ctx *Constant_LOCAL_IDContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTSqlParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTSqlParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTime_zone is called when production time_zone is entered.
func (s *BaseTSqlParserListener) EnterTime_zone(ctx *Time_zoneContext) {}

// ExitTime_zone is called when production time_zone is exited.
func (s *BaseTSqlParserListener) ExitTime_zone(ctx *Time_zoneContext) {}

// EnterPrimitive_expression is called when production primitive_expression is entered.
func (s *BaseTSqlParserListener) EnterPrimitive_expression(ctx *Primitive_expressionContext) {}

// ExitPrimitive_expression is called when production primitive_expression is exited.
func (s *BaseTSqlParserListener) ExitPrimitive_expression(ctx *Primitive_expressionContext) {}

// EnterCase_expression is called when production case_expression is entered.
func (s *BaseTSqlParserListener) EnterCase_expression(ctx *Case_expressionContext) {}

// ExitCase_expression is called when production case_expression is exited.
func (s *BaseTSqlParserListener) ExitCase_expression(ctx *Case_expressionContext) {}

// EnterUnary_operator_expression is called when production unary_operator_expression is entered.
func (s *BaseTSqlParserListener) EnterUnary_operator_expression(ctx *Unary_operator_expressionContext) {
}

// ExitUnary_operator_expression is called when production unary_operator_expression is exited.
func (s *BaseTSqlParserListener) ExitUnary_operator_expression(ctx *Unary_operator_expressionContext) {
}

// EnterBracket_expression is called when production bracket_expression is entered.
func (s *BaseTSqlParserListener) EnterBracket_expression(ctx *Bracket_expressionContext) {}

// ExitBracket_expression is called when production bracket_expression is exited.
func (s *BaseTSqlParserListener) ExitBracket_expression(ctx *Bracket_expressionContext) {}

// EnterConstant_expression is called when production constant_expression is entered.
func (s *BaseTSqlParserListener) EnterConstant_expression(ctx *Constant_expressionContext) {}

// ExitConstant_expression is called when production constant_expression is exited.
func (s *BaseTSqlParserListener) ExitConstant_expression(ctx *Constant_expressionContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseTSqlParserListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseTSqlParserListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterWith_expression is called when production with_expression is entered.
func (s *BaseTSqlParserListener) EnterWith_expression(ctx *With_expressionContext) {}

// ExitWith_expression is called when production with_expression is exited.
func (s *BaseTSqlParserListener) ExitWith_expression(ctx *With_expressionContext) {}

// EnterCommon_table_expression is called when production common_table_expression is entered.
func (s *BaseTSqlParserListener) EnterCommon_table_expression(ctx *Common_table_expressionContext) {}

// ExitCommon_table_expression is called when production common_table_expression is exited.
func (s *BaseTSqlParserListener) ExitCommon_table_expression(ctx *Common_table_expressionContext) {}

// EnterUpdate_elem is called when production update_elem is entered.
func (s *BaseTSqlParserListener) EnterUpdate_elem(ctx *Update_elemContext) {}

// ExitUpdate_elem is called when production update_elem is exited.
func (s *BaseTSqlParserListener) ExitUpdate_elem(ctx *Update_elemContext) {}

// EnterUpdate_elem_merge is called when production update_elem_merge is entered.
func (s *BaseTSqlParserListener) EnterUpdate_elem_merge(ctx *Update_elem_mergeContext) {}

// ExitUpdate_elem_merge is called when production update_elem_merge is exited.
func (s *BaseTSqlParserListener) ExitUpdate_elem_merge(ctx *Update_elem_mergeContext) {}

// EnterSearch_condition is called when production search_condition is entered.
func (s *BaseTSqlParserListener) EnterSearch_condition(ctx *Search_conditionContext) {}

// ExitSearch_condition is called when production search_condition is exited.
func (s *BaseTSqlParserListener) ExitSearch_condition(ctx *Search_conditionContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseTSqlParserListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseTSqlParserListener) ExitPredicate(ctx *PredicateContext) {}

// EnterQuery_expression is called when production query_expression is entered.
func (s *BaseTSqlParserListener) EnterQuery_expression(ctx *Query_expressionContext) {}

// ExitQuery_expression is called when production query_expression is exited.
func (s *BaseTSqlParserListener) ExitQuery_expression(ctx *Query_expressionContext) {}

// EnterSql_union is called when production sql_union is entered.
func (s *BaseTSqlParserListener) EnterSql_union(ctx *Sql_unionContext) {}

// ExitSql_union is called when production sql_union is exited.
func (s *BaseTSqlParserListener) ExitSql_union(ctx *Sql_unionContext) {}

// EnterQuery_specification is called when production query_specification is entered.
func (s *BaseTSqlParserListener) EnterQuery_specification(ctx *Query_specificationContext) {}

// ExitQuery_specification is called when production query_specification is exited.
func (s *BaseTSqlParserListener) ExitQuery_specification(ctx *Query_specificationContext) {}

// EnterTop_clause is called when production top_clause is entered.
func (s *BaseTSqlParserListener) EnterTop_clause(ctx *Top_clauseContext) {}

// ExitTop_clause is called when production top_clause is exited.
func (s *BaseTSqlParserListener) ExitTop_clause(ctx *Top_clauseContext) {}

// EnterTop_percent is called when production top_percent is entered.
func (s *BaseTSqlParserListener) EnterTop_percent(ctx *Top_percentContext) {}

// ExitTop_percent is called when production top_percent is exited.
func (s *BaseTSqlParserListener) ExitTop_percent(ctx *Top_percentContext) {}

// EnterTop_count is called when production top_count is entered.
func (s *BaseTSqlParserListener) EnterTop_count(ctx *Top_countContext) {}

// ExitTop_count is called when production top_count is exited.
func (s *BaseTSqlParserListener) ExitTop_count(ctx *Top_countContext) {}

// EnterOrder_by_clause is called when production order_by_clause is entered.
func (s *BaseTSqlParserListener) EnterOrder_by_clause(ctx *Order_by_clauseContext) {}

// ExitOrder_by_clause is called when production order_by_clause is exited.
func (s *BaseTSqlParserListener) ExitOrder_by_clause(ctx *Order_by_clauseContext) {}

// EnterFor_clause is called when production for_clause is entered.
func (s *BaseTSqlParserListener) EnterFor_clause(ctx *For_clauseContext) {}

// ExitFor_clause is called when production for_clause is exited.
func (s *BaseTSqlParserListener) ExitFor_clause(ctx *For_clauseContext) {}

// EnterXml_common_directives is called when production xml_common_directives is entered.
func (s *BaseTSqlParserListener) EnterXml_common_directives(ctx *Xml_common_directivesContext) {}

// ExitXml_common_directives is called when production xml_common_directives is exited.
func (s *BaseTSqlParserListener) ExitXml_common_directives(ctx *Xml_common_directivesContext) {}

// EnterOrder_by_expression is called when production order_by_expression is entered.
func (s *BaseTSqlParserListener) EnterOrder_by_expression(ctx *Order_by_expressionContext) {}

// ExitOrder_by_expression is called when production order_by_expression is exited.
func (s *BaseTSqlParserListener) ExitOrder_by_expression(ctx *Order_by_expressionContext) {}

// EnterGroup_by_item is called when production group_by_item is entered.
func (s *BaseTSqlParserListener) EnterGroup_by_item(ctx *Group_by_itemContext) {}

// ExitGroup_by_item is called when production group_by_item is exited.
func (s *BaseTSqlParserListener) ExitGroup_by_item(ctx *Group_by_itemContext) {}

// EnterOption_clause is called when production option_clause is entered.
func (s *BaseTSqlParserListener) EnterOption_clause(ctx *Option_clauseContext) {}

// ExitOption_clause is called when production option_clause is exited.
func (s *BaseTSqlParserListener) ExitOption_clause(ctx *Option_clauseContext) {}

// EnterOption is called when production option is entered.
func (s *BaseTSqlParserListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseTSqlParserListener) ExitOption(ctx *OptionContext) {}

// EnterOptimize_for_arg is called when production optimize_for_arg is entered.
func (s *BaseTSqlParserListener) EnterOptimize_for_arg(ctx *Optimize_for_argContext) {}

// ExitOptimize_for_arg is called when production optimize_for_arg is exited.
func (s *BaseTSqlParserListener) ExitOptimize_for_arg(ctx *Optimize_for_argContext) {}

// EnterSelect_list is called when production select_list is entered.
func (s *BaseTSqlParserListener) EnterSelect_list(ctx *Select_listContext) {}

// ExitSelect_list is called when production select_list is exited.
func (s *BaseTSqlParserListener) ExitSelect_list(ctx *Select_listContext) {}

// EnterUdt_method_arguments is called when production udt_method_arguments is entered.
func (s *BaseTSqlParserListener) EnterUdt_method_arguments(ctx *Udt_method_argumentsContext) {}

// ExitUdt_method_arguments is called when production udt_method_arguments is exited.
func (s *BaseTSqlParserListener) ExitUdt_method_arguments(ctx *Udt_method_argumentsContext) {}

// EnterAsterisk is called when production asterisk is entered.
func (s *BaseTSqlParserListener) EnterAsterisk(ctx *AsteriskContext) {}

// ExitAsterisk is called when production asterisk is exited.
func (s *BaseTSqlParserListener) ExitAsterisk(ctx *AsteriskContext) {}

// EnterColumn_elem is called when production column_elem is entered.
func (s *BaseTSqlParserListener) EnterColumn_elem(ctx *Column_elemContext) {}

// ExitColumn_elem is called when production column_elem is exited.
func (s *BaseTSqlParserListener) ExitColumn_elem(ctx *Column_elemContext) {}

// EnterUdt_elem is called when production udt_elem is entered.
func (s *BaseTSqlParserListener) EnterUdt_elem(ctx *Udt_elemContext) {}

// ExitUdt_elem is called when production udt_elem is exited.
func (s *BaseTSqlParserListener) ExitUdt_elem(ctx *Udt_elemContext) {}

// EnterExpression_elem is called when production expression_elem is entered.
func (s *BaseTSqlParserListener) EnterExpression_elem(ctx *Expression_elemContext) {}

// ExitExpression_elem is called when production expression_elem is exited.
func (s *BaseTSqlParserListener) ExitExpression_elem(ctx *Expression_elemContext) {}

// EnterSelect_list_elem is called when production select_list_elem is entered.
func (s *BaseTSqlParserListener) EnterSelect_list_elem(ctx *Select_list_elemContext) {}

// ExitSelect_list_elem is called when production select_list_elem is exited.
func (s *BaseTSqlParserListener) ExitSelect_list_elem(ctx *Select_list_elemContext) {}

// EnterTable_sources is called when production table_sources is entered.
func (s *BaseTSqlParserListener) EnterTable_sources(ctx *Table_sourcesContext) {}

// ExitTable_sources is called when production table_sources is exited.
func (s *BaseTSqlParserListener) ExitTable_sources(ctx *Table_sourcesContext) {}

// EnterTable_source is called when production table_source is entered.
func (s *BaseTSqlParserListener) EnterTable_source(ctx *Table_sourceContext) {}

// ExitTable_source is called when production table_source is exited.
func (s *BaseTSqlParserListener) ExitTable_source(ctx *Table_sourceContext) {}

// EnterTable_source_item_joined is called when production table_source_item_joined is entered.
func (s *BaseTSqlParserListener) EnterTable_source_item_joined(ctx *Table_source_item_joinedContext) {
}

// ExitTable_source_item_joined is called when production table_source_item_joined is exited.
func (s *BaseTSqlParserListener) ExitTable_source_item_joined(ctx *Table_source_item_joinedContext) {}

// EnterTable_source_item is called when production table_source_item is entered.
func (s *BaseTSqlParserListener) EnterTable_source_item(ctx *Table_source_itemContext) {}

// ExitTable_source_item is called when production table_source_item is exited.
func (s *BaseTSqlParserListener) ExitTable_source_item(ctx *Table_source_itemContext) {}

// EnterOpen_xml is called when production open_xml is entered.
func (s *BaseTSqlParserListener) EnterOpen_xml(ctx *Open_xmlContext) {}

// ExitOpen_xml is called when production open_xml is exited.
func (s *BaseTSqlParserListener) ExitOpen_xml(ctx *Open_xmlContext) {}

// EnterOpen_json is called when production open_json is entered.
func (s *BaseTSqlParserListener) EnterOpen_json(ctx *Open_jsonContext) {}

// ExitOpen_json is called when production open_json is exited.
func (s *BaseTSqlParserListener) ExitOpen_json(ctx *Open_jsonContext) {}

// EnterJson_declaration is called when production json_declaration is entered.
func (s *BaseTSqlParserListener) EnterJson_declaration(ctx *Json_declarationContext) {}

// ExitJson_declaration is called when production json_declaration is exited.
func (s *BaseTSqlParserListener) ExitJson_declaration(ctx *Json_declarationContext) {}

// EnterJson_column_declaration is called when production json_column_declaration is entered.
func (s *BaseTSqlParserListener) EnterJson_column_declaration(ctx *Json_column_declarationContext) {}

// ExitJson_column_declaration is called when production json_column_declaration is exited.
func (s *BaseTSqlParserListener) ExitJson_column_declaration(ctx *Json_column_declarationContext) {}

// EnterSchema_declaration is called when production schema_declaration is entered.
func (s *BaseTSqlParserListener) EnterSchema_declaration(ctx *Schema_declarationContext) {}

// ExitSchema_declaration is called when production schema_declaration is exited.
func (s *BaseTSqlParserListener) ExitSchema_declaration(ctx *Schema_declarationContext) {}

// EnterColumn_declaration is called when production column_declaration is entered.
func (s *BaseTSqlParserListener) EnterColumn_declaration(ctx *Column_declarationContext) {}

// ExitColumn_declaration is called when production column_declaration is exited.
func (s *BaseTSqlParserListener) ExitColumn_declaration(ctx *Column_declarationContext) {}

// EnterChange_table is called when production change_table is entered.
func (s *BaseTSqlParserListener) EnterChange_table(ctx *Change_tableContext) {}

// ExitChange_table is called when production change_table is exited.
func (s *BaseTSqlParserListener) ExitChange_table(ctx *Change_tableContext) {}

// EnterChange_table_changes is called when production change_table_changes is entered.
func (s *BaseTSqlParserListener) EnterChange_table_changes(ctx *Change_table_changesContext) {}

// ExitChange_table_changes is called when production change_table_changes is exited.
func (s *BaseTSqlParserListener) ExitChange_table_changes(ctx *Change_table_changesContext) {}

// EnterChange_table_version is called when production change_table_version is entered.
func (s *BaseTSqlParserListener) EnterChange_table_version(ctx *Change_table_versionContext) {}

// ExitChange_table_version is called when production change_table_version is exited.
func (s *BaseTSqlParserListener) ExitChange_table_version(ctx *Change_table_versionContext) {}

// EnterJoin_part is called when production join_part is entered.
func (s *BaseTSqlParserListener) EnterJoin_part(ctx *Join_partContext) {}

// ExitJoin_part is called when production join_part is exited.
func (s *BaseTSqlParserListener) ExitJoin_part(ctx *Join_partContext) {}

// EnterJoin_on is called when production join_on is entered.
func (s *BaseTSqlParserListener) EnterJoin_on(ctx *Join_onContext) {}

// ExitJoin_on is called when production join_on is exited.
func (s *BaseTSqlParserListener) ExitJoin_on(ctx *Join_onContext) {}

// EnterCross_join is called when production cross_join is entered.
func (s *BaseTSqlParserListener) EnterCross_join(ctx *Cross_joinContext) {}

// ExitCross_join is called when production cross_join is exited.
func (s *BaseTSqlParserListener) ExitCross_join(ctx *Cross_joinContext) {}

// EnterApply_ is called when production apply_ is entered.
func (s *BaseTSqlParserListener) EnterApply_(ctx *Apply_Context) {}

// ExitApply_ is called when production apply_ is exited.
func (s *BaseTSqlParserListener) ExitApply_(ctx *Apply_Context) {}

// EnterPivot is called when production pivot is entered.
func (s *BaseTSqlParserListener) EnterPivot(ctx *PivotContext) {}

// ExitPivot is called when production pivot is exited.
func (s *BaseTSqlParserListener) ExitPivot(ctx *PivotContext) {}

// EnterUnpivot is called when production unpivot is entered.
func (s *BaseTSqlParserListener) EnterUnpivot(ctx *UnpivotContext) {}

// ExitUnpivot is called when production unpivot is exited.
func (s *BaseTSqlParserListener) ExitUnpivot(ctx *UnpivotContext) {}

// EnterPivot_clause is called when production pivot_clause is entered.
func (s *BaseTSqlParserListener) EnterPivot_clause(ctx *Pivot_clauseContext) {}

// ExitPivot_clause is called when production pivot_clause is exited.
func (s *BaseTSqlParserListener) ExitPivot_clause(ctx *Pivot_clauseContext) {}

// EnterUnpivot_clause is called when production unpivot_clause is entered.
func (s *BaseTSqlParserListener) EnterUnpivot_clause(ctx *Unpivot_clauseContext) {}

// ExitUnpivot_clause is called when production unpivot_clause is exited.
func (s *BaseTSqlParserListener) ExitUnpivot_clause(ctx *Unpivot_clauseContext) {}

// EnterFull_column_name_list is called when production full_column_name_list is entered.
func (s *BaseTSqlParserListener) EnterFull_column_name_list(ctx *Full_column_name_listContext) {}

// ExitFull_column_name_list is called when production full_column_name_list is exited.
func (s *BaseTSqlParserListener) ExitFull_column_name_list(ctx *Full_column_name_listContext) {}

// EnterTable_name_with_hint is called when production table_name_with_hint is entered.
func (s *BaseTSqlParserListener) EnterTable_name_with_hint(ctx *Table_name_with_hintContext) {}

// ExitTable_name_with_hint is called when production table_name_with_hint is exited.
func (s *BaseTSqlParserListener) ExitTable_name_with_hint(ctx *Table_name_with_hintContext) {}

// EnterRowset_function is called when production rowset_function is entered.
func (s *BaseTSqlParserListener) EnterRowset_function(ctx *Rowset_functionContext) {}

// ExitRowset_function is called when production rowset_function is exited.
func (s *BaseTSqlParserListener) ExitRowset_function(ctx *Rowset_functionContext) {}

// EnterBulk_option is called when production bulk_option is entered.
func (s *BaseTSqlParserListener) EnterBulk_option(ctx *Bulk_optionContext) {}

// ExitBulk_option is called when production bulk_option is exited.
func (s *BaseTSqlParserListener) ExitBulk_option(ctx *Bulk_optionContext) {}

// EnterDerived_table is called when production derived_table is entered.
func (s *BaseTSqlParserListener) EnterDerived_table(ctx *Derived_tableContext) {}

// ExitDerived_table is called when production derived_table is exited.
func (s *BaseTSqlParserListener) ExitDerived_table(ctx *Derived_tableContext) {}

// EnterRANKING_WINDOWED_FUNC is called when production RANKING_WINDOWED_FUNC is entered.
func (s *BaseTSqlParserListener) EnterRANKING_WINDOWED_FUNC(ctx *RANKING_WINDOWED_FUNCContext) {}

// ExitRANKING_WINDOWED_FUNC is called when production RANKING_WINDOWED_FUNC is exited.
func (s *BaseTSqlParserListener) ExitRANKING_WINDOWED_FUNC(ctx *RANKING_WINDOWED_FUNCContext) {}

// EnterAGGREGATE_WINDOWED_FUNC is called when production AGGREGATE_WINDOWED_FUNC is entered.
func (s *BaseTSqlParserListener) EnterAGGREGATE_WINDOWED_FUNC(ctx *AGGREGATE_WINDOWED_FUNCContext) {}

// ExitAGGREGATE_WINDOWED_FUNC is called when production AGGREGATE_WINDOWED_FUNC is exited.
func (s *BaseTSqlParserListener) ExitAGGREGATE_WINDOWED_FUNC(ctx *AGGREGATE_WINDOWED_FUNCContext) {}

// EnterANALYTIC_WINDOWED_FUNC is called when production ANALYTIC_WINDOWED_FUNC is entered.
func (s *BaseTSqlParserListener) EnterANALYTIC_WINDOWED_FUNC(ctx *ANALYTIC_WINDOWED_FUNCContext) {}

// ExitANALYTIC_WINDOWED_FUNC is called when production ANALYTIC_WINDOWED_FUNC is exited.
func (s *BaseTSqlParserListener) ExitANALYTIC_WINDOWED_FUNC(ctx *ANALYTIC_WINDOWED_FUNCContext) {}

// EnterBUILT_IN_FUNC is called when production BUILT_IN_FUNC is entered.
func (s *BaseTSqlParserListener) EnterBUILT_IN_FUNC(ctx *BUILT_IN_FUNCContext) {}

// ExitBUILT_IN_FUNC is called when production BUILT_IN_FUNC is exited.
func (s *BaseTSqlParserListener) ExitBUILT_IN_FUNC(ctx *BUILT_IN_FUNCContext) {}

// EnterSCALAR_FUNCTION is called when production SCALAR_FUNCTION is entered.
func (s *BaseTSqlParserListener) EnterSCALAR_FUNCTION(ctx *SCALAR_FUNCTIONContext) {}

// ExitSCALAR_FUNCTION is called when production SCALAR_FUNCTION is exited.
func (s *BaseTSqlParserListener) ExitSCALAR_FUNCTION(ctx *SCALAR_FUNCTIONContext) {}

// EnterFREE_TEXT is called when production FREE_TEXT is entered.
func (s *BaseTSqlParserListener) EnterFREE_TEXT(ctx *FREE_TEXTContext) {}

// ExitFREE_TEXT is called when production FREE_TEXT is exited.
func (s *BaseTSqlParserListener) ExitFREE_TEXT(ctx *FREE_TEXTContext) {}

// EnterPARTITION_FUNC is called when production PARTITION_FUNC is entered.
func (s *BaseTSqlParserListener) EnterPARTITION_FUNC(ctx *PARTITION_FUNCContext) {}

// ExitPARTITION_FUNC is called when production PARTITION_FUNC is exited.
func (s *BaseTSqlParserListener) ExitPARTITION_FUNC(ctx *PARTITION_FUNCContext) {}

// EnterPartition_function is called when production partition_function is entered.
func (s *BaseTSqlParserListener) EnterPartition_function(ctx *Partition_functionContext) {}

// ExitPartition_function is called when production partition_function is exited.
func (s *BaseTSqlParserListener) ExitPartition_function(ctx *Partition_functionContext) {}

// EnterFreetext_function is called when production freetext_function is entered.
func (s *BaseTSqlParserListener) EnterFreetext_function(ctx *Freetext_functionContext) {}

// ExitFreetext_function is called when production freetext_function is exited.
func (s *BaseTSqlParserListener) ExitFreetext_function(ctx *Freetext_functionContext) {}

// EnterFreetext_predicate is called when production freetext_predicate is entered.
func (s *BaseTSqlParserListener) EnterFreetext_predicate(ctx *Freetext_predicateContext) {}

// ExitFreetext_predicate is called when production freetext_predicate is exited.
func (s *BaseTSqlParserListener) ExitFreetext_predicate(ctx *Freetext_predicateContext) {}

// EnterBINARY_CHECKSUM is called when production BINARY_CHECKSUM is entered.
func (s *BaseTSqlParserListener) EnterBINARY_CHECKSUM(ctx *BINARY_CHECKSUMContext) {}

// ExitBINARY_CHECKSUM is called when production BINARY_CHECKSUM is exited.
func (s *BaseTSqlParserListener) ExitBINARY_CHECKSUM(ctx *BINARY_CHECKSUMContext) {}

// EnterCAST is called when production CAST is entered.
func (s *BaseTSqlParserListener) EnterCAST(ctx *CASTContext) {}

// ExitCAST is called when production CAST is exited.
func (s *BaseTSqlParserListener) ExitCAST(ctx *CASTContext) {}

// EnterTRY_CAST is called when production TRY_CAST is entered.
func (s *BaseTSqlParserListener) EnterTRY_CAST(ctx *TRY_CASTContext) {}

// ExitTRY_CAST is called when production TRY_CAST is exited.
func (s *BaseTSqlParserListener) ExitTRY_CAST(ctx *TRY_CASTContext) {}

// EnterCONVERT is called when production CONVERT is entered.
func (s *BaseTSqlParserListener) EnterCONVERT(ctx *CONVERTContext) {}

// ExitCONVERT is called when production CONVERT is exited.
func (s *BaseTSqlParserListener) ExitCONVERT(ctx *CONVERTContext) {}

// EnterCHECKSUM is called when production CHECKSUM is entered.
func (s *BaseTSqlParserListener) EnterCHECKSUM(ctx *CHECKSUMContext) {}

// ExitCHECKSUM is called when production CHECKSUM is exited.
func (s *BaseTSqlParserListener) ExitCHECKSUM(ctx *CHECKSUMContext) {}

// EnterCOALESCE is called when production COALESCE is entered.
func (s *BaseTSqlParserListener) EnterCOALESCE(ctx *COALESCEContext) {}

// ExitCOALESCE is called when production COALESCE is exited.
func (s *BaseTSqlParserListener) ExitCOALESCE(ctx *COALESCEContext) {}

// EnterCURRENT_TIMESTAMP is called when production CURRENT_TIMESTAMP is entered.
func (s *BaseTSqlParserListener) EnterCURRENT_TIMESTAMP(ctx *CURRENT_TIMESTAMPContext) {}

// ExitCURRENT_TIMESTAMP is called when production CURRENT_TIMESTAMP is exited.
func (s *BaseTSqlParserListener) ExitCURRENT_TIMESTAMP(ctx *CURRENT_TIMESTAMPContext) {}

// EnterCURRENT_USER is called when production CURRENT_USER is entered.
func (s *BaseTSqlParserListener) EnterCURRENT_USER(ctx *CURRENT_USERContext) {}

// ExitCURRENT_USER is called when production CURRENT_USER is exited.
func (s *BaseTSqlParserListener) ExitCURRENT_USER(ctx *CURRENT_USERContext) {}

// EnterDATEADD is called when production DATEADD is entered.
func (s *BaseTSqlParserListener) EnterDATEADD(ctx *DATEADDContext) {}

// ExitDATEADD is called when production DATEADD is exited.
func (s *BaseTSqlParserListener) ExitDATEADD(ctx *DATEADDContext) {}

// EnterDATEDIFF is called when production DATEDIFF is entered.
func (s *BaseTSqlParserListener) EnterDATEDIFF(ctx *DATEDIFFContext) {}

// ExitDATEDIFF is called when production DATEDIFF is exited.
func (s *BaseTSqlParserListener) ExitDATEDIFF(ctx *DATEDIFFContext) {}

// EnterDATENAME is called when production DATENAME is entered.
func (s *BaseTSqlParserListener) EnterDATENAME(ctx *DATENAMEContext) {}

// ExitDATENAME is called when production DATENAME is exited.
func (s *BaseTSqlParserListener) ExitDATENAME(ctx *DATENAMEContext) {}

// EnterDATEPART is called when production DATEPART is entered.
func (s *BaseTSqlParserListener) EnterDATEPART(ctx *DATEPARTContext) {}

// ExitDATEPART is called when production DATEPART is exited.
func (s *BaseTSqlParserListener) ExitDATEPART(ctx *DATEPARTContext) {}

// EnterGETDATE is called when production GETDATE is entered.
func (s *BaseTSqlParserListener) EnterGETDATE(ctx *GETDATEContext) {}

// ExitGETDATE is called when production GETDATE is exited.
func (s *BaseTSqlParserListener) ExitGETDATE(ctx *GETDATEContext) {}

// EnterGETUTCDATE is called when production GETUTCDATE is entered.
func (s *BaseTSqlParserListener) EnterGETUTCDATE(ctx *GETUTCDATEContext) {}

// ExitGETUTCDATE is called when production GETUTCDATE is exited.
func (s *BaseTSqlParserListener) ExitGETUTCDATE(ctx *GETUTCDATEContext) {}

// EnterIDENTITY is called when production IDENTITY is entered.
func (s *BaseTSqlParserListener) EnterIDENTITY(ctx *IDENTITYContext) {}

// ExitIDENTITY is called when production IDENTITY is exited.
func (s *BaseTSqlParserListener) ExitIDENTITY(ctx *IDENTITYContext) {}

// EnterMIN_ACTIVE_ROWVERSION is called when production MIN_ACTIVE_ROWVERSION is entered.
func (s *BaseTSqlParserListener) EnterMIN_ACTIVE_ROWVERSION(ctx *MIN_ACTIVE_ROWVERSIONContext) {}

// ExitMIN_ACTIVE_ROWVERSION is called when production MIN_ACTIVE_ROWVERSION is exited.
func (s *BaseTSqlParserListener) ExitMIN_ACTIVE_ROWVERSION(ctx *MIN_ACTIVE_ROWVERSIONContext) {}

// EnterNULLIF is called when production NULLIF is entered.
func (s *BaseTSqlParserListener) EnterNULLIF(ctx *NULLIFContext) {}

// ExitNULLIF is called when production NULLIF is exited.
func (s *BaseTSqlParserListener) ExitNULLIF(ctx *NULLIFContext) {}

// EnterSTUFF is called when production STUFF is entered.
func (s *BaseTSqlParserListener) EnterSTUFF(ctx *STUFFContext) {}

// ExitSTUFF is called when production STUFF is exited.
func (s *BaseTSqlParserListener) ExitSTUFF(ctx *STUFFContext) {}

// EnterSESSION_USER is called when production SESSION_USER is entered.
func (s *BaseTSqlParserListener) EnterSESSION_USER(ctx *SESSION_USERContext) {}

// ExitSESSION_USER is called when production SESSION_USER is exited.
func (s *BaseTSqlParserListener) ExitSESSION_USER(ctx *SESSION_USERContext) {}

// EnterSYSTEM_USER is called when production SYSTEM_USER is entered.
func (s *BaseTSqlParserListener) EnterSYSTEM_USER(ctx *SYSTEM_USERContext) {}

// ExitSYSTEM_USER is called when production SYSTEM_USER is exited.
func (s *BaseTSqlParserListener) ExitSYSTEM_USER(ctx *SYSTEM_USERContext) {}

// EnterUSER is called when production USER is entered.
func (s *BaseTSqlParserListener) EnterUSER(ctx *USERContext) {}

// ExitUSER is called when production USER is exited.
func (s *BaseTSqlParserListener) ExitUSER(ctx *USERContext) {}

// EnterISNULL is called when production ISNULL is entered.
func (s *BaseTSqlParserListener) EnterISNULL(ctx *ISNULLContext) {}

// ExitISNULL is called when production ISNULL is exited.
func (s *BaseTSqlParserListener) ExitISNULL(ctx *ISNULLContext) {}

// EnterXML_DATA_TYPE_FUNC is called when production XML_DATA_TYPE_FUNC is entered.
func (s *BaseTSqlParserListener) EnterXML_DATA_TYPE_FUNC(ctx *XML_DATA_TYPE_FUNCContext) {}

// ExitXML_DATA_TYPE_FUNC is called when production XML_DATA_TYPE_FUNC is exited.
func (s *BaseTSqlParserListener) ExitXML_DATA_TYPE_FUNC(ctx *XML_DATA_TYPE_FUNCContext) {}

// EnterIIF is called when production IIF is entered.
func (s *BaseTSqlParserListener) EnterIIF(ctx *IIFContext) {}

// ExitIIF is called when production IIF is exited.
func (s *BaseTSqlParserListener) ExitIIF(ctx *IIFContext) {}

// EnterSTRINGAGG is called when production STRINGAGG is entered.
func (s *BaseTSqlParserListener) EnterSTRINGAGG(ctx *STRINGAGGContext) {}

// ExitSTRINGAGG is called when production STRINGAGG is exited.
func (s *BaseTSqlParserListener) ExitSTRINGAGG(ctx *STRINGAGGContext) {}

// EnterXml_data_type_methods is called when production xml_data_type_methods is entered.
func (s *BaseTSqlParserListener) EnterXml_data_type_methods(ctx *Xml_data_type_methodsContext) {}

// ExitXml_data_type_methods is called when production xml_data_type_methods is exited.
func (s *BaseTSqlParserListener) ExitXml_data_type_methods(ctx *Xml_data_type_methodsContext) {}

// EnterValue_method is called when production value_method is entered.
func (s *BaseTSqlParserListener) EnterValue_method(ctx *Value_methodContext) {}

// ExitValue_method is called when production value_method is exited.
func (s *BaseTSqlParserListener) ExitValue_method(ctx *Value_methodContext) {}

// EnterValue_call is called when production value_call is entered.
func (s *BaseTSqlParserListener) EnterValue_call(ctx *Value_callContext) {}

// ExitValue_call is called when production value_call is exited.
func (s *BaseTSqlParserListener) ExitValue_call(ctx *Value_callContext) {}

// EnterQuery_method is called when production query_method is entered.
func (s *BaseTSqlParserListener) EnterQuery_method(ctx *Query_methodContext) {}

// ExitQuery_method is called when production query_method is exited.
func (s *BaseTSqlParserListener) ExitQuery_method(ctx *Query_methodContext) {}

// EnterQuery_call is called when production query_call is entered.
func (s *BaseTSqlParserListener) EnterQuery_call(ctx *Query_callContext) {}

// ExitQuery_call is called when production query_call is exited.
func (s *BaseTSqlParserListener) ExitQuery_call(ctx *Query_callContext) {}

// EnterExist_method is called when production exist_method is entered.
func (s *BaseTSqlParserListener) EnterExist_method(ctx *Exist_methodContext) {}

// ExitExist_method is called when production exist_method is exited.
func (s *BaseTSqlParserListener) ExitExist_method(ctx *Exist_methodContext) {}

// EnterExist_call is called when production exist_call is entered.
func (s *BaseTSqlParserListener) EnterExist_call(ctx *Exist_callContext) {}

// ExitExist_call is called when production exist_call is exited.
func (s *BaseTSqlParserListener) ExitExist_call(ctx *Exist_callContext) {}

// EnterModify_method is called when production modify_method is entered.
func (s *BaseTSqlParserListener) EnterModify_method(ctx *Modify_methodContext) {}

// ExitModify_method is called when production modify_method is exited.
func (s *BaseTSqlParserListener) ExitModify_method(ctx *Modify_methodContext) {}

// EnterModify_call is called when production modify_call is entered.
func (s *BaseTSqlParserListener) EnterModify_call(ctx *Modify_callContext) {}

// ExitModify_call is called when production modify_call is exited.
func (s *BaseTSqlParserListener) ExitModify_call(ctx *Modify_callContext) {}

// EnterNodes_method is called when production nodes_method is entered.
func (s *BaseTSqlParserListener) EnterNodes_method(ctx *Nodes_methodContext) {}

// ExitNodes_method is called when production nodes_method is exited.
func (s *BaseTSqlParserListener) ExitNodes_method(ctx *Nodes_methodContext) {}

// EnterSwitch_section is called when production switch_section is entered.
func (s *BaseTSqlParserListener) EnterSwitch_section(ctx *Switch_sectionContext) {}

// ExitSwitch_section is called when production switch_section is exited.
func (s *BaseTSqlParserListener) ExitSwitch_section(ctx *Switch_sectionContext) {}

// EnterSwitch_search_condition_section is called when production switch_search_condition_section is entered.
func (s *BaseTSqlParserListener) EnterSwitch_search_condition_section(ctx *Switch_search_condition_sectionContext) {
}

// ExitSwitch_search_condition_section is called when production switch_search_condition_section is exited.
func (s *BaseTSqlParserListener) ExitSwitch_search_condition_section(ctx *Switch_search_condition_sectionContext) {
}

// EnterAs_column_alias is called when production as_column_alias is entered.
func (s *BaseTSqlParserListener) EnterAs_column_alias(ctx *As_column_aliasContext) {}

// ExitAs_column_alias is called when production as_column_alias is exited.
func (s *BaseTSqlParserListener) ExitAs_column_alias(ctx *As_column_aliasContext) {}

// EnterAs_table_alias is called when production as_table_alias is entered.
func (s *BaseTSqlParserListener) EnterAs_table_alias(ctx *As_table_aliasContext) {}

// ExitAs_table_alias is called when production as_table_alias is exited.
func (s *BaseTSqlParserListener) ExitAs_table_alias(ctx *As_table_aliasContext) {}

// EnterTable_alias is called when production table_alias is entered.
func (s *BaseTSqlParserListener) EnterTable_alias(ctx *Table_aliasContext) {}

// ExitTable_alias is called when production table_alias is exited.
func (s *BaseTSqlParserListener) ExitTable_alias(ctx *Table_aliasContext) {}

// EnterWith_table_hints is called when production with_table_hints is entered.
func (s *BaseTSqlParserListener) EnterWith_table_hints(ctx *With_table_hintsContext) {}

// ExitWith_table_hints is called when production with_table_hints is exited.
func (s *BaseTSqlParserListener) ExitWith_table_hints(ctx *With_table_hintsContext) {}

// EnterInsert_with_table_hints is called when production insert_with_table_hints is entered.
func (s *BaseTSqlParserListener) EnterInsert_with_table_hints(ctx *Insert_with_table_hintsContext) {}

// ExitInsert_with_table_hints is called when production insert_with_table_hints is exited.
func (s *BaseTSqlParserListener) ExitInsert_with_table_hints(ctx *Insert_with_table_hintsContext) {}

// EnterTable_hint is called when production table_hint is entered.
func (s *BaseTSqlParserListener) EnterTable_hint(ctx *Table_hintContext) {}

// ExitTable_hint is called when production table_hint is exited.
func (s *BaseTSqlParserListener) ExitTable_hint(ctx *Table_hintContext) {}

// EnterIndex_value is called when production index_value is entered.
func (s *BaseTSqlParserListener) EnterIndex_value(ctx *Index_valueContext) {}

// ExitIndex_value is called when production index_value is exited.
func (s *BaseTSqlParserListener) ExitIndex_value(ctx *Index_valueContext) {}

// EnterColumn_alias_list is called when production column_alias_list is entered.
func (s *BaseTSqlParserListener) EnterColumn_alias_list(ctx *Column_alias_listContext) {}

// ExitColumn_alias_list is called when production column_alias_list is exited.
func (s *BaseTSqlParserListener) ExitColumn_alias_list(ctx *Column_alias_listContext) {}

// EnterColumn_alias is called when production column_alias is entered.
func (s *BaseTSqlParserListener) EnterColumn_alias(ctx *Column_aliasContext) {}

// ExitColumn_alias is called when production column_alias is exited.
func (s *BaseTSqlParserListener) ExitColumn_alias(ctx *Column_aliasContext) {}

// EnterTable_value_constructor is called when production table_value_constructor is entered.
func (s *BaseTSqlParserListener) EnterTable_value_constructor(ctx *Table_value_constructorContext) {}

// ExitTable_value_constructor is called when production table_value_constructor is exited.
func (s *BaseTSqlParserListener) ExitTable_value_constructor(ctx *Table_value_constructorContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BaseTSqlParserListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BaseTSqlParserListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterRanking_windowed_function is called when production ranking_windowed_function is entered.
func (s *BaseTSqlParserListener) EnterRanking_windowed_function(ctx *Ranking_windowed_functionContext) {
}

// ExitRanking_windowed_function is called when production ranking_windowed_function is exited.
func (s *BaseTSqlParserListener) ExitRanking_windowed_function(ctx *Ranking_windowed_functionContext) {
}

// EnterAggregate_windowed_function is called when production aggregate_windowed_function is entered.
func (s *BaseTSqlParserListener) EnterAggregate_windowed_function(ctx *Aggregate_windowed_functionContext) {
}

// ExitAggregate_windowed_function is called when production aggregate_windowed_function is exited.
func (s *BaseTSqlParserListener) ExitAggregate_windowed_function(ctx *Aggregate_windowed_functionContext) {
}

// EnterAnalytic_windowed_function is called when production analytic_windowed_function is entered.
func (s *BaseTSqlParserListener) EnterAnalytic_windowed_function(ctx *Analytic_windowed_functionContext) {
}

// ExitAnalytic_windowed_function is called when production analytic_windowed_function is exited.
func (s *BaseTSqlParserListener) ExitAnalytic_windowed_function(ctx *Analytic_windowed_functionContext) {
}

// EnterAll_distinct_expression is called when production all_distinct_expression is entered.
func (s *BaseTSqlParserListener) EnterAll_distinct_expression(ctx *All_distinct_expressionContext) {}

// ExitAll_distinct_expression is called when production all_distinct_expression is exited.
func (s *BaseTSqlParserListener) ExitAll_distinct_expression(ctx *All_distinct_expressionContext) {}

// EnterOver_clause is called when production over_clause is entered.
func (s *BaseTSqlParserListener) EnterOver_clause(ctx *Over_clauseContext) {}

// ExitOver_clause is called when production over_clause is exited.
func (s *BaseTSqlParserListener) ExitOver_clause(ctx *Over_clauseContext) {}

// EnterRow_or_range_clause is called when production row_or_range_clause is entered.
func (s *BaseTSqlParserListener) EnterRow_or_range_clause(ctx *Row_or_range_clauseContext) {}

// ExitRow_or_range_clause is called when production row_or_range_clause is exited.
func (s *BaseTSqlParserListener) ExitRow_or_range_clause(ctx *Row_or_range_clauseContext) {}

// EnterWindow_frame_extent is called when production window_frame_extent is entered.
func (s *BaseTSqlParserListener) EnterWindow_frame_extent(ctx *Window_frame_extentContext) {}

// ExitWindow_frame_extent is called when production window_frame_extent is exited.
func (s *BaseTSqlParserListener) ExitWindow_frame_extent(ctx *Window_frame_extentContext) {}

// EnterWindow_frame_bound is called when production window_frame_bound is entered.
func (s *BaseTSqlParserListener) EnterWindow_frame_bound(ctx *Window_frame_boundContext) {}

// ExitWindow_frame_bound is called when production window_frame_bound is exited.
func (s *BaseTSqlParserListener) ExitWindow_frame_bound(ctx *Window_frame_boundContext) {}

// EnterWindow_frame_preceding is called when production window_frame_preceding is entered.
func (s *BaseTSqlParserListener) EnterWindow_frame_preceding(ctx *Window_frame_precedingContext) {}

// ExitWindow_frame_preceding is called when production window_frame_preceding is exited.
func (s *BaseTSqlParserListener) ExitWindow_frame_preceding(ctx *Window_frame_precedingContext) {}

// EnterWindow_frame_following is called when production window_frame_following is entered.
func (s *BaseTSqlParserListener) EnterWindow_frame_following(ctx *Window_frame_followingContext) {}

// ExitWindow_frame_following is called when production window_frame_following is exited.
func (s *BaseTSqlParserListener) ExitWindow_frame_following(ctx *Window_frame_followingContext) {}

// EnterCreate_database_option is called when production create_database_option is entered.
func (s *BaseTSqlParserListener) EnterCreate_database_option(ctx *Create_database_optionContext) {}

// ExitCreate_database_option is called when production create_database_option is exited.
func (s *BaseTSqlParserListener) ExitCreate_database_option(ctx *Create_database_optionContext) {}

// EnterDatabase_filestream_option is called when production database_filestream_option is entered.
func (s *BaseTSqlParserListener) EnterDatabase_filestream_option(ctx *Database_filestream_optionContext) {
}

// ExitDatabase_filestream_option is called when production database_filestream_option is exited.
func (s *BaseTSqlParserListener) ExitDatabase_filestream_option(ctx *Database_filestream_optionContext) {
}

// EnterDatabase_file_spec is called when production database_file_spec is entered.
func (s *BaseTSqlParserListener) EnterDatabase_file_spec(ctx *Database_file_specContext) {}

// ExitDatabase_file_spec is called when production database_file_spec is exited.
func (s *BaseTSqlParserListener) ExitDatabase_file_spec(ctx *Database_file_specContext) {}

// EnterFile_group is called when production file_group is entered.
func (s *BaseTSqlParserListener) EnterFile_group(ctx *File_groupContext) {}

// ExitFile_group is called when production file_group is exited.
func (s *BaseTSqlParserListener) ExitFile_group(ctx *File_groupContext) {}

// EnterFile_spec is called when production file_spec is entered.
func (s *BaseTSqlParserListener) EnterFile_spec(ctx *File_specContext) {}

// ExitFile_spec is called when production file_spec is exited.
func (s *BaseTSqlParserListener) ExitFile_spec(ctx *File_specContext) {}

// EnterEntity_name is called when production entity_name is entered.
func (s *BaseTSqlParserListener) EnterEntity_name(ctx *Entity_nameContext) {}

// ExitEntity_name is called when production entity_name is exited.
func (s *BaseTSqlParserListener) ExitEntity_name(ctx *Entity_nameContext) {}

// EnterEntity_name_for_azure_dw is called when production entity_name_for_azure_dw is entered.
func (s *BaseTSqlParserListener) EnterEntity_name_for_azure_dw(ctx *Entity_name_for_azure_dwContext) {
}

// ExitEntity_name_for_azure_dw is called when production entity_name_for_azure_dw is exited.
func (s *BaseTSqlParserListener) ExitEntity_name_for_azure_dw(ctx *Entity_name_for_azure_dwContext) {}

// EnterEntity_name_for_parallel_dw is called when production entity_name_for_parallel_dw is entered.
func (s *BaseTSqlParserListener) EnterEntity_name_for_parallel_dw(ctx *Entity_name_for_parallel_dwContext) {
}

// ExitEntity_name_for_parallel_dw is called when production entity_name_for_parallel_dw is exited.
func (s *BaseTSqlParserListener) ExitEntity_name_for_parallel_dw(ctx *Entity_name_for_parallel_dwContext) {
}

// EnterFull_table_name is called when production full_table_name is entered.
func (s *BaseTSqlParserListener) EnterFull_table_name(ctx *Full_table_nameContext) {}

// ExitFull_table_name is called when production full_table_name is exited.
func (s *BaseTSqlParserListener) ExitFull_table_name(ctx *Full_table_nameContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseTSqlParserListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseTSqlParserListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterSimple_name is called when production simple_name is entered.
func (s *BaseTSqlParserListener) EnterSimple_name(ctx *Simple_nameContext) {}

// ExitSimple_name is called when production simple_name is exited.
func (s *BaseTSqlParserListener) ExitSimple_name(ctx *Simple_nameContext) {}

// EnterFunc_proc_name_schema is called when production func_proc_name_schema is entered.
func (s *BaseTSqlParserListener) EnterFunc_proc_name_schema(ctx *Func_proc_name_schemaContext) {}

// ExitFunc_proc_name_schema is called when production func_proc_name_schema is exited.
func (s *BaseTSqlParserListener) ExitFunc_proc_name_schema(ctx *Func_proc_name_schemaContext) {}

// EnterFunc_proc_name_database_schema is called when production func_proc_name_database_schema is entered.
func (s *BaseTSqlParserListener) EnterFunc_proc_name_database_schema(ctx *Func_proc_name_database_schemaContext) {
}

// ExitFunc_proc_name_database_schema is called when production func_proc_name_database_schema is exited.
func (s *BaseTSqlParserListener) ExitFunc_proc_name_database_schema(ctx *Func_proc_name_database_schemaContext) {
}

// EnterFunc_proc_name_server_database_schema is called when production func_proc_name_server_database_schema is entered.
func (s *BaseTSqlParserListener) EnterFunc_proc_name_server_database_schema(ctx *Func_proc_name_server_database_schemaContext) {
}

// ExitFunc_proc_name_server_database_schema is called when production func_proc_name_server_database_schema is exited.
func (s *BaseTSqlParserListener) ExitFunc_proc_name_server_database_schema(ctx *Func_proc_name_server_database_schemaContext) {
}

// EnterDdl_object is called when production ddl_object is entered.
func (s *BaseTSqlParserListener) EnterDdl_object(ctx *Ddl_objectContext) {}

// ExitDdl_object is called when production ddl_object is exited.
func (s *BaseTSqlParserListener) ExitDdl_object(ctx *Ddl_objectContext) {}

// EnterFull_column_name is called when production full_column_name is entered.
func (s *BaseTSqlParserListener) EnterFull_column_name(ctx *Full_column_nameContext) {}

// ExitFull_column_name is called when production full_column_name is exited.
func (s *BaseTSqlParserListener) ExitFull_column_name(ctx *Full_column_nameContext) {}

// EnterColumn_name_list_with_order is called when production column_name_list_with_order is entered.
func (s *BaseTSqlParserListener) EnterColumn_name_list_with_order(ctx *Column_name_list_with_orderContext) {
}

// ExitColumn_name_list_with_order is called when production column_name_list_with_order is exited.
func (s *BaseTSqlParserListener) ExitColumn_name_list_with_order(ctx *Column_name_list_with_orderContext) {
}

// EnterInsert_column_name_list is called when production insert_column_name_list is entered.
func (s *BaseTSqlParserListener) EnterInsert_column_name_list(ctx *Insert_column_name_listContext) {}

// ExitInsert_column_name_list is called when production insert_column_name_list is exited.
func (s *BaseTSqlParserListener) ExitInsert_column_name_list(ctx *Insert_column_name_listContext) {}

// EnterInsert_column_id is called when production insert_column_id is entered.
func (s *BaseTSqlParserListener) EnterInsert_column_id(ctx *Insert_column_idContext) {}

// ExitInsert_column_id is called when production insert_column_id is exited.
func (s *BaseTSqlParserListener) ExitInsert_column_id(ctx *Insert_column_idContext) {}

// EnterColumn_name_list is called when production column_name_list is entered.
func (s *BaseTSqlParserListener) EnterColumn_name_list(ctx *Column_name_listContext) {}

// ExitColumn_name_list is called when production column_name_list is exited.
func (s *BaseTSqlParserListener) ExitColumn_name_list(ctx *Column_name_listContext) {}

// EnterCursor_name is called when production cursor_name is entered.
func (s *BaseTSqlParserListener) EnterCursor_name(ctx *Cursor_nameContext) {}

// ExitCursor_name is called when production cursor_name is exited.
func (s *BaseTSqlParserListener) ExitCursor_name(ctx *Cursor_nameContext) {}

// EnterOn_off is called when production on_off is entered.
func (s *BaseTSqlParserListener) EnterOn_off(ctx *On_offContext) {}

// ExitOn_off is called when production on_off is exited.
func (s *BaseTSqlParserListener) ExitOn_off(ctx *On_offContext) {}

// EnterClustered is called when production clustered is entered.
func (s *BaseTSqlParserListener) EnterClustered(ctx *ClusteredContext) {}

// ExitClustered is called when production clustered is exited.
func (s *BaseTSqlParserListener) ExitClustered(ctx *ClusteredContext) {}

// EnterNull_notnull is called when production null_notnull is entered.
func (s *BaseTSqlParserListener) EnterNull_notnull(ctx *Null_notnullContext) {}

// ExitNull_notnull is called when production null_notnull is exited.
func (s *BaseTSqlParserListener) ExitNull_notnull(ctx *Null_notnullContext) {}

// EnterNull_or_default is called when production null_or_default is entered.
func (s *BaseTSqlParserListener) EnterNull_or_default(ctx *Null_or_defaultContext) {}

// ExitNull_or_default is called when production null_or_default is exited.
func (s *BaseTSqlParserListener) ExitNull_or_default(ctx *Null_or_defaultContext) {}

// EnterScalar_function_name is called when production scalar_function_name is entered.
func (s *BaseTSqlParserListener) EnterScalar_function_name(ctx *Scalar_function_nameContext) {}

// ExitScalar_function_name is called when production scalar_function_name is exited.
func (s *BaseTSqlParserListener) ExitScalar_function_name(ctx *Scalar_function_nameContext) {}

// EnterBegin_conversation_timer is called when production begin_conversation_timer is entered.
func (s *BaseTSqlParserListener) EnterBegin_conversation_timer(ctx *Begin_conversation_timerContext) {
}

// ExitBegin_conversation_timer is called when production begin_conversation_timer is exited.
func (s *BaseTSqlParserListener) ExitBegin_conversation_timer(ctx *Begin_conversation_timerContext) {}

// EnterBegin_conversation_dialog is called when production begin_conversation_dialog is entered.
func (s *BaseTSqlParserListener) EnterBegin_conversation_dialog(ctx *Begin_conversation_dialogContext) {
}

// ExitBegin_conversation_dialog is called when production begin_conversation_dialog is exited.
func (s *BaseTSqlParserListener) ExitBegin_conversation_dialog(ctx *Begin_conversation_dialogContext) {
}

// EnterContract_name is called when production contract_name is entered.
func (s *BaseTSqlParserListener) EnterContract_name(ctx *Contract_nameContext) {}

// ExitContract_name is called when production contract_name is exited.
func (s *BaseTSqlParserListener) ExitContract_name(ctx *Contract_nameContext) {}

// EnterService_name is called when production service_name is entered.
func (s *BaseTSqlParserListener) EnterService_name(ctx *Service_nameContext) {}

// ExitService_name is called when production service_name is exited.
func (s *BaseTSqlParserListener) ExitService_name(ctx *Service_nameContext) {}

// EnterEnd_conversation is called when production end_conversation is entered.
func (s *BaseTSqlParserListener) EnterEnd_conversation(ctx *End_conversationContext) {}

// ExitEnd_conversation is called when production end_conversation is exited.
func (s *BaseTSqlParserListener) ExitEnd_conversation(ctx *End_conversationContext) {}

// EnterWaitfor_conversation is called when production waitfor_conversation is entered.
func (s *BaseTSqlParserListener) EnterWaitfor_conversation(ctx *Waitfor_conversationContext) {}

// ExitWaitfor_conversation is called when production waitfor_conversation is exited.
func (s *BaseTSqlParserListener) ExitWaitfor_conversation(ctx *Waitfor_conversationContext) {}

// EnterGet_conversation is called when production get_conversation is entered.
func (s *BaseTSqlParserListener) EnterGet_conversation(ctx *Get_conversationContext) {}

// ExitGet_conversation is called when production get_conversation is exited.
func (s *BaseTSqlParserListener) ExitGet_conversation(ctx *Get_conversationContext) {}

// EnterQueue_id is called when production queue_id is entered.
func (s *BaseTSqlParserListener) EnterQueue_id(ctx *Queue_idContext) {}

// ExitQueue_id is called when production queue_id is exited.
func (s *BaseTSqlParserListener) ExitQueue_id(ctx *Queue_idContext) {}

// EnterSend_conversation is called when production send_conversation is entered.
func (s *BaseTSqlParserListener) EnterSend_conversation(ctx *Send_conversationContext) {}

// ExitSend_conversation is called when production send_conversation is exited.
func (s *BaseTSqlParserListener) ExitSend_conversation(ctx *Send_conversationContext) {}

// EnterData_type is called when production data_type is entered.
func (s *BaseTSqlParserListener) EnterData_type(ctx *Data_typeContext) {}

// ExitData_type is called when production data_type is exited.
func (s *BaseTSqlParserListener) ExitData_type(ctx *Data_typeContext) {}

// EnterDefault_value is called when production default_value is entered.
func (s *BaseTSqlParserListener) EnterDefault_value(ctx *Default_valueContext) {}

// ExitDefault_value is called when production default_value is exited.
func (s *BaseTSqlParserListener) ExitDefault_value(ctx *Default_valueContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseTSqlParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseTSqlParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterSign is called when production sign is entered.
func (s *BaseTSqlParserListener) EnterSign(ctx *SignContext) {}

// ExitSign is called when production sign is exited.
func (s *BaseTSqlParserListener) ExitSign(ctx *SignContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseTSqlParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseTSqlParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterId_ is called when production id_ is entered.
func (s *BaseTSqlParserListener) EnterId_(ctx *Id_Context) {}

// ExitId_ is called when production id_ is exited.
func (s *BaseTSqlParserListener) ExitId_(ctx *Id_Context) {}

// EnterSimple_id is called when production simple_id is entered.
func (s *BaseTSqlParserListener) EnterSimple_id(ctx *Simple_idContext) {}

// ExitSimple_id is called when production simple_id is exited.
func (s *BaseTSqlParserListener) ExitSimple_id(ctx *Simple_idContext) {}

// EnterId_or_string is called when production id_or_string is entered.
func (s *BaseTSqlParserListener) EnterId_or_string(ctx *Id_or_stringContext) {}

// ExitId_or_string is called when production id_or_string is exited.
func (s *BaseTSqlParserListener) ExitId_or_string(ctx *Id_or_stringContext) {}

// EnterComparison_operator is called when production comparison_operator is entered.
func (s *BaseTSqlParserListener) EnterComparison_operator(ctx *Comparison_operatorContext) {}

// ExitComparison_operator is called when production comparison_operator is exited.
func (s *BaseTSqlParserListener) ExitComparison_operator(ctx *Comparison_operatorContext) {}

// EnterAssignment_operator is called when production assignment_operator is entered.
func (s *BaseTSqlParserListener) EnterAssignment_operator(ctx *Assignment_operatorContext) {}

// ExitAssignment_operator is called when production assignment_operator is exited.
func (s *BaseTSqlParserListener) ExitAssignment_operator(ctx *Assignment_operatorContext) {}

// EnterFile_size is called when production file_size is entered.
func (s *BaseTSqlParserListener) EnterFile_size(ctx *File_sizeContext) {}

// ExitFile_size is called when production file_size is exited.
func (s *BaseTSqlParserListener) ExitFile_size(ctx *File_sizeContext) {}
