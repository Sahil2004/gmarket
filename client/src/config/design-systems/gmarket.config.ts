import { IDesignSystem } from '../design-system';

export class GMarket implements IDesignSystem {
  app: IDesignSystem['app'] = {
    APP_NAME: 'G-Market',
  };
  surface: IDesignSystem['surface'] = {
    MatFormFieldAppearance: 'outline',
    MatCardAppearance: 'outlined',
  };
}
