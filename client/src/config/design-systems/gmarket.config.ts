import { IDesignSystem } from '../../types/design-system.types';

export class GMarket implements IDesignSystem {
  app: IDesignSystem['app'] = {
    APP_NAME: 'G-Market',
  };
  color: IDesignSystem['color'] = {
    theme: 'system',
  };
  typography: IDesignSystem['typography'] = {
    subtitleTextWeight: 300,
  };
  surface: IDesignSystem['surface'] = {
    MatFormFieldAppearance: 'outline',
    MatCardAppearance: 'outlined',
  };
  devConfig: IDesignSystem['devConfig'] = {
    throttlingTimeMs: 2 * 1000, // 2 seconds
  };
}
