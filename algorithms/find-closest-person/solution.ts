function findClosest(x: number, y: number, z: number): number {
  const xDistance = x > z ? x - z : z - x;
  const yDistance = y > z ? y - z : z - y;
  if (xDistance < yDistance) {
    return 1;
  }
  if (xDistance > yDistance) {
    return 2;
  }
  return 0;
}
