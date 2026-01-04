import { InjectionToken } from '@angular/core';
import type { IDesignSystem } from '../types/design-system.types';
import { designSystemFactory } from './design-system.factory';

export const DESIGN_SYSTEM = new InjectionToken<IDesignSystem>('DESIGN_SYSTEM', {
  providedIn: 'root',
  factory: designSystemFactory,
});
