import { createSelector, createFeatureSelector } from '@ngrx/store';
import { InfraRoleEntityState, infraRoleEntityAdapter } from './infra-role.reducer';
import { routeParams } from 'app/route.selectors';
import { find } from 'lodash/fp';

export const infraRoleState = createFeatureSelector<InfraRoleEntityState>('infraRoles');
export const {
  selectAll: allInfraRoles,
  selectEntities: roleEntities
} = infraRoleEntityAdapter.getSelectors(infraRoleState);

export const infraRoleStatus = createSelector(
  infraRoleState,
  (state) => state.rolesStatus
);

export const getAllStatus = createSelector(
  infraRoleState,
  (state) => state.getAllStatus
);

export const infraRoleFromRoute = createSelector(
  roleEntities,
  routeParams,
  (state, { name }) => find({ name }, state)
);
