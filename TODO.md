Todo list for esctl-go
====
## Repo 
- [x] Add Cobra to repo for subcommands
- [x] Separate authentication into its own package
- [x] Separate helper utilities into their own package
- [x] Separate elastic functions into their own package
- [ ] Add ability to switch between elastic clusters
- [ ] Add automatic documentation for command usage (https://godoc.org/github.com/spf13/cobra/doc)
- [ ] Make top level command `esctl` display avaiable subcommands if no
  arguments are present
- [ ] Make second level command `esctl get` display available subcommands if no
  arguements are present
- [ ] Determine which verbs `esctl` should use other than `get`
- [ ] Add login command to generate a token
- [ ] Add config file support

## Function porting from https://github.com/slmingol/escli/blob/master/es_funcs.bash
- [ ] escli_ls
- [ ] escli_lsl
- [ ] usage_chk1
- [ ] usage_chk2
- [ ] usage_chk3
- [ ] usage_chk4
- [ ] usage_chk5
- [ ] list_nodes
- [ ] list_nodes_storage
- [ ] show_shards
- [ ] show_big_shards
- [ ] show_small_shards
- [ ] relo_s
- [ ] cancel_relo_shard
- [ ] cancel_relo_shards_all
- [ ] increase_balance_throttle
- [ ] reset_balance_throttle
- [ ] show_balance_throttle
- [ ] show_recovery
- [ ] show_recovery_full
- [ ] enable_readonly_idx_pattern
- [ ] disable_readonly_idx_pattern
- [ ] enable_readonly_idxs
- [ ] disable_readonly_idxs
- [ ] show_readonly_idxs_full
- [ ] show_readonly_idxs
- [ ] estop
- [ ] estop_recovery
- [ ] estop_relo
- [ ] show_health
- [ ] show_watermarks
- [ ] show_state
- [ ] showcfg_cluster
- [ ] showcfg_num_shards_per_idx
- [ ] showcfg_shard_allocations
- [ ] explain_allocations
- [ ] help_cat
- [ ] help_indices
- [ ] show_idx_sizes
- [ ] show_idx_stats
- [ ] delete_idx
- [ ] exclude_node_name
- [ ] show_excluded_nodes
- [ ] clear_excluded_nodes