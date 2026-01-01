import { InjectionToken } from '@angular/core';
import type { IDesignSystem } from './design-system';
import { designSystemFactory } from './design-system.factory';

export const DESIGN_SYSTEM = new InjectionToken<IDesignSystem>('DESIGN_SYSTEM', {
  providedIn: 'root',
  factory: designSystemFactory,
});
