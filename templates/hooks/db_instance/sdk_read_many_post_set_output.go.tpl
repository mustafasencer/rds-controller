	if ko.Status.ACKResourceMetadata != nil && ko.Status.ACKResourceMetadata.ARN != nil {
        resourceARN := (*string)(ko.Status.ACKResourceMetadata.ARN)
        tags, err := rm.getTags(ctx, *resourceARN)
        if err != nil {
            return nil, err
        }
        ko.Spec.Tags = tags
	}
	if !instanceAvailable(&resource{ko}) {
		// Setting resource synced condition to false will trigger a requeue of
		// the resource. No need to return a requeue error here.
		ackcondition.SetSynced(&resource{ko}, corev1.ConditionFalse, nil, nil)
	}
	if len(r.ko.Spec.VPCSecurityGroupIDs) > 0 {
		// If the desired resource has security groups specified then update the spec of the latest resource with the
		// security groups from the status. This is done so that when an instance is created without security groups
		// and gets a default security group attached to it, it is not overwritten with no security groups from the
		// desired resource.
		sgIDs := make([]*string, len(ko.Status.VPCSecurityGroups))
		for i, sg := range ko.Status.VPCSecurityGroups {
			id := *sg.VPCSecurityGroupID
			sgIDs[i] = &id
		}
		ko.Spec.VPCSecurityGroupIDs = sgIDs
	}
