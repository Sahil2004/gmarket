import { IDesignSystem } from '../design-system';

export class GMarket implements IDesignSystem {
  surface: IDesignSystem['surface'] = {
    MatFormFieldAppearance: 'outline',
    MatCardAppearance: 'outlined',
  };
}
