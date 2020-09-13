package go_repo_gen

func GetNotNilList(effects []DBEffect, DB Interpreter) ([]Any, error) {
	var resList []Any

	for _, effect := range effects {
		ctx := effect.Apply(DB)
		// break if error
		if !ctx.OK {
			return resList, ctx.Err
		}
		// append result to list
		if ctx.Ctx != nil {
			resList = append(resList, ctx.Ctx)
		}
	}

	return resList, nil
}
