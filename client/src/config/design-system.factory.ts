import { IDesignSystem } from '../types/design-system.types';
import { environment } from '../environments/environment';

import { GMarket } from './design-systems/gmarket.config';

export function designSystemFactory(): IDesignSystem {
  switch (environment.APP_NAME) {
    case 'gmarket':
      return new GMarket();
    default:
      throw new Error(`Unknown design system for APP_NAME: ${environment.APP_NAME}`);
  }
}
