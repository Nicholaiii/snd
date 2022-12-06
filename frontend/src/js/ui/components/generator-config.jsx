import Types from '/js/ui/components/generator/types';
import { Input } from '/js/ui/components/index';

export default () => {
	return {
		view(vnode) {
			return (
				<div className='lh-copy'>
					<Input
						label='Random Seed'
						placeholder='121FA0GA...'
						value={vnode.attrs.value.seed}
						oninput={(e) => vnode.attrs.onchange('seed', e.target.value)}
					></Input>{' '}
					<div className='btn btn-primary' onclick={() => vnode.attrs.onchange('seed', Math.ceil(Math.random() * 1000000000))}>
						Reroll
					</div>
					<div className='mb2 o-70'>
						A random seed will make the random number generator deterministic. The same seed will result in the same generated random
						values.
					</div>
					<div className='divider' />
					{vnode.attrs.config.map((val, i) => {
						return (
							<div>
								{m(Types[val.type].view, {
									value: vnode.attrs.value[val.key],
									oninput: (newVal) => {
										vnode.attrs.onchange(val.key, newVal);
									},
									inEdit: false,
									label: val.name,
								})}
								<span className='o-70 mt2 mb2'>{val.description}</span>
								<div className='divider' />
							</div>
						);
					})}
				</div>
			);
		},
	};
};