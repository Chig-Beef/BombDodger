# Bomb Dodger
This is a port of an old Python program that has many agents that learn to dodge bombs.
Each generation contains a certain number of agents, and as the game plays, they must survive dropping bombs.
The agents that survive longer have greater chance of reproduction, and so the next generation should survive longer.
Bombs, however, drop more frequently as time goes on, so there is a cap to how long a given agent can live.

## Agents
Agents have brain with these inputs:
1. Is there a bomb above?
2. Is there a bomb to the left?
3. Is there a bomb to the right?
4. Is there a wall to the left?
5. Is there a wall to the right?

Agents can then use this information in their neural network to manage these outputs:
1. Move left.
2. Move right.

The smart agents learn to move out of the way of upcoming bombs, while not trapping themselves into a wall.
This behaviour is very simple, and can even be chanced upon within a single generation.

## Future
This simple demonstration simply shows that given enough brain power agents can learn simple tasks.
From this, slightly more complicated simulations could be attempted, such as:
1. The agents can catch certain objects, which give more fitness points.
2. 2D movement.
3. Obstacles and jumping.
4. Inter-agent strategy.
5. Ageing/degrading mechanisms.
