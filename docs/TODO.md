Todo list for esctl-go
## Repo 
- [x] Add Cobra to repo for subcommands
- [x] Separate authentication into its own package
- [x] Separate helper utilities into their own package
- [x] Separate elastic functions into their own package
- [x] Add CI for building packages
- [x] Add common badges to README
- [x] Add ability to switch between elastic clusters
  * [x] Via command
  * [x] Via flag
- [ ] Add automatic documentation for command usage (https://godoc.org/github.com/spf13/cobra/doc)
- [x] Make top level command `esctl` display avaiable subcommands if no
  arguments are present
- [x] Make second level command `esctl get` display available subcommands if no
  arguements are present
- [ ] Determine which verbs `esctl` should use other than `get`
- [x] Add support for optional modules at compile time
  * [x] Create a script to initialize new modules
  * [x] Create a Makefile that allows compilation of optional modules
- [x] Add the ability to debug (requires a code change to activate)
- [ ] Mock [Watcher](https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api.html) apis
  - [ ] Put Watcher
  - [x] Get Watcher
  - [ ] Delete Watcher
  - [ ] Execute Watcher
  - [x] Get Watcher Stats
  - [ ] Ack Watcher
  - [x] Activate Watcher
  - [x] Deactivate Watcher
  - [ ] Stop Watcher Service
  - [ ] Start Watcher Service
  - [x] List Watchers
  - [x] Show Active Watchers
  - [x] Show Inactive Watchers
- [x] Add support for output format flag
- [ ] Mock [Cluster](https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster.html) apis
- [x] Improve error handling to be more verbose (somewhat complete, needs
  improvement)
- [x] Add API module to make requests in VERB ENDPOINT form to support not yet implemented endpoints
   * [x] MVP
   * [x] Support for output format flag and other existing flags
   * [x] General alignment with esapi.Request objects (duplicate code)
   * [x] Rename functions/structs/interfaces. (Probably GenericRequest and
     similar)
   * [x] ~~Move api subcommand under escmd.~~ It shouldn't be considered an
     extension/option. Every subcommand should probably have its own package.
     Keeping the current structure
- [ ] Mock functions from [kubectl
  config](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#config)
   * [x] current-context
   * [ ] delete-cluster
   * [ ] delete-context
   * [ ] get-clusters
   * [x] get-contexts
   * [ ] rename-context
   * [ ] set
   * [ ] set-cluster
   * [ ] set-context
   * [ ] set-credentials
   * [x] use-context
   * [x] view (currently read)
   * [x] generate
   * [x] validate
- [x] add ability to execute artibrary commands rather than store
  username/password in config file


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
